package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"myChat-API/internal/service"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.Password == "" {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := h.As.Signup(ctx, req.Username, req.Password); err != nil {
		log.Println(err)
		http.Error(w, "Failed to signup.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.FormValue("username")
	password := r.FormValue("password")

	token, err := h.As.Login(ctx, username, password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to login.", http.StatusInternalServerError)
		return
	}

	res := GetTokenResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func trimHeaderBearer(tokenStr string) (string, error) {
	headers := strings.Split(tokenStr, " ")

	if len(headers) != 2 {
		return "", errors.New("Invalid token header format")
	}

	bearerStr, accessToken := headers[0], headers[1]
	if bearerStr != "Bearer" || accessToken == "" {
		return "", errors.New("Invalid token header format")
	}

	return accessToken, nil
}

func validateJWT(r *http.Request) (map[string]interface{}, error) {

	tokenStr := r.Header.Get("Authorization")
	trimedToken, err := trimHeaderBearer(tokenStr)
	if err != nil {
		return nil, err
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(trimedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []bytemy_secret_key")
		return []byte(service.SecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Token parse error: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, err
}

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := validateJWT(r)

		username, ok := claims["username"]
		if !ok {
			http.Error(w, "Authorization failed. Invalid token", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		// uuid is interface{} type.
		// so, uuid needs to be asserted.
		r = SetUsername(r, username.(string))

		next.ServeHTTP(w, r)
	})
}
