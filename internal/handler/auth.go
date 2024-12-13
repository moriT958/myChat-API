package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"myChat-API/internal/common"
	"myChat-API/internal/model"
	"myChat-API/internal/schema"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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
		return []byte("sample_secret_key"), nil
	})
	if err != nil {
		log.Printf("Failed to parse token: %+v\n", err)
		return nil, fmt.Errorf("Token parse error: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, err

}

func generateJWT(username string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString([]byte("sample_secret_key"))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req schema.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.Password == "" {
		log.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	u, err := createUser(req.Username, req.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	if err := h.DAO.SaveUser(u); err != nil {
		log.Println(err)
		http.Error(w, "Failed to Save User", http.StatusInternalServerError)
		return
	}

	res := schema.CreateUserResponse{
		Uuid:     u.Uuid.String(),
		CreateAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func createUser(name string, password string) (model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	u := model.User{
		Uuid:      uuid.New(),
		Username:  name,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}
	return u, nil
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to generate password hash", http.StatusInternalServerError)
		return
	}

	u, err := h.DAO.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		http.Error(w, "User doesn't exists.", http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password)); err != nil {
		http.Error(w, "Invalid password. Auth failed.", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	token, err := generateJWT(u.Username)
	res := schema.GetTokenResponse{
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
		r = common.SetUsername(r, username.(string))

		next.ServeHTTP(w, r)
	})
}
