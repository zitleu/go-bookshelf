package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID         uuid.UUID
	Username     string
	Email        string
	PasswordHash string `json:"-"`
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

func (u *User) ToPublic() UserPublic {
	return UserPublic{
		UUID:      u.UUID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) ToSummary() UserSummary {
	return UserSummary{
		ID:       u.UUID,
		Username: u.Username,
	}
}

type UserPublic struct {
	UUID      uuid.UUID
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserSummary struct {
	ID       uuid.UUID
	Username string
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken string     `json:"access_token"`
	TokenType   string     `json:"token_type"`
	ExpiresIn   int        `json:"expires_in"`
	User        UserPublic `json:"user"`
}

type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
}
