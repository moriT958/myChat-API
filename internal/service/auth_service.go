package service

import (
	"context"
	"errors"
	"fmt"
	"myChat-API2/internal/domain"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// TODO:
// Add secret key to config obj.

type IAuthService interface {
	Signup(context.Context, string, string) (string, error)
	Login(context.Context, string, string) (string, error)
	SeeUserDetail(context.Context, string) (domain.User, string, error)
}

type AuthService struct {
	UserRepo domain.IUserRepository
}

func NewAuthService(ur domain.IUserRepository) *AuthService {
	return &AuthService{UserRepo: ur}
}

func (s *AuthService) Signup(ctx context.Context, username string, password string) (string, error) {
	if password == "" {
		return "", errors.New("require password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	u := domain.User{
		ID:       uuid.NewString(),
		Name:     username,
		Password: string(hash),
	}
	if err := s.UserRepo.Save(ctx, u); err != nil {
		return "", err
	}

	return u.ID, nil
}

func generateJWT(userId string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *AuthService) Login(ctx context.Context, username string, password string) (string, error) {
	// Find user by username.
	u, err := s.UserRepo.GetByName(ctx, username)
	if err != nil {
		return "", fmt.Errorf("User doesn't exit: %w", err)
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
	token, err := generateJWT(u.ID)
	return token, nil
}

func (s *AuthService) SeeUserDetail(ctx context.Context, userID string) (domain.User, string, error) {
	user, err := s.UserRepo.GetByID(ctx, userID)
	if err != nil {
		return domain.User{}, "", err
	}

	createdAt, err := s.UserRepo.GetCreatedAtByID(ctx, userID)
	if err != nil {
		return domain.User{}, "", err
	}

	return user, createdAt, nil
}
