package service

import (
	"context"
	"fmt"
	"myChat-API/internal/domain"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Ur domain.UserRepositorier
}

func NewAuthService(ur domain.UserRepositorier) *AuthService {
	return &AuthService{Ur: ur}
}

func (s *AuthService) Signup(ctx context.Context, username string, password string) error {
	// Create user
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := domain.User{
		Uuid:     uuid.NewString(),
		Username: username,
		Password: string(hash),
	}
	if err := s.Ur.SaveUser(ctx, u); err != nil {
		return err
	}

	return nil
}

// TODO:
// Add secret key to config obj.
var SecretKey = os.Getenv("SECRET_KEY")

func generateJWT(username string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *AuthService) Login(ctx context.Context, username string, password string) (string, error) {
	// Find user by username.
	u, err := s.Ur.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// Hash password.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to generate password hash: %w", err)
	}

	// Check if password is correct.
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return "", err
	}

	// Generate jwt token.
	token, err := generateJWT(u.Uuid)
	return token, nil
}
