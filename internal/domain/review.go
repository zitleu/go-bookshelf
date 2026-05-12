package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID
	BookID    string
	UserID    string
	Rating    int
	Title     sql.NullString
	Content   string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type ReviewResponse struct {
	ID        uuid.UUID   `json:"id"`
	BookID    string      `json:"book_id"`
	UserID    string      `json:"user_id"`
	Rating    int         `json:"rating"`
	Title     *string     `json:"title"`
	Content   string      `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
	User      UserSummary `json:"user"`
}

type CreateReviewRequest struct {
	ID        uuid.UUID   `json:"id"`
	BookID    string      `json:"book_id"`
	UserID    string      `json:"user_id"`
	Rating    int         `json:"rating"`
	Title     *string     `json:"title"`
	Content   string      `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
	User      UserSummary `json:"user"`
}

type UpdateReviewRequest struct {
	ID        uuid.UUID  `json:"id"`
	BookID    *string    `json:"book_id"`
	UserID    *string    `json:"user_id"`
	Rating    *int       `json:"rating"`
	Title     *string    `json:"title"`
	Content   *string    `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ReviewListResponse struct {
	Data       []ReviewResponse
	Pagination Pagination
}

func (r *Review) ToResponse(user *User) ReviewResponse
