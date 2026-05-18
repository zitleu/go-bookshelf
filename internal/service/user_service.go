package service

import (
	"context"
	"errors"

	"github.com/bookshelf/monolith/internal/domain"
	"github.com/bookshelf/monolith/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
	ErrUsernameExists     = errors.New("username already exists")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidUsername    = errors.New("invalid username")
	ErrInvalidEmail       = errors.New("invalid email")

	ErrUsernameValidation = errors.New("username too short")
	ErrPasswordValidation = errors.New("password too short")
	ErrEmailValidation    = errors.New("email cannot be empty")
)

type UserService struct {
	repo      *repository.UserRepository
	jwtSecret string
}

func (s *UserService) Register(ctx context.Context, req domain.RegisterRequest) (*domain.AuthResponse, error) {
	// валидация
	if len(req.Username) < 4 {
		return nil, ErrUsernameValidation
	}

	if len(req.Password) < 8 {
		return nil, ErrPasswordValidation
	}

	if len(req.Email) == 0 {
		return nil, ErrEmailValidation
	}

	exists, err := s.repo.EmailExists(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrUserExists
	}

	exists, err = s.repo.UsernameExists(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	err = s.repo.Create(ctx, &user)
	if err != nil {
		return nil, err
	}

	authResponse := domain.AuthResponse{
		AccessToken: "",
		TokenType:   "",
		ExpiresIn:   0,
		User: domain.UserPublic{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return &authResponse, nil
}
