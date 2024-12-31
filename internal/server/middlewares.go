package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (h *TodoServer) AuthMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := validateJWT(r)

		userID, ok := claims["id"]
		if !ok {
			http.Error(w, "Authorization failed. Invalid token", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		// userID is interface{} type.
		// so, userID needs to be asserted.
		r = setUserID(r, userID.(string))

		next.ServeHTTP(w, r)
	})
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
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Token parse error: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, err
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
