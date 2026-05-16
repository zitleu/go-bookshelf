package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID      `db:"id"`
	BookID    uuid.UUID      `db:"book_id"`
	UserID    uuid.UUID      `db:"user_id"`
	Rating    int            `db:"rating"`
	Title     sql.NullString `db:"title"`
	Content   string         `db:"content"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type ReviewResponse struct {
	ID        uuid.UUID   `json:"id"`
	BookID    string      `json:"book_id"`
	UserID    string      `json:"user_id"`
	Rating    int         `json:"rating"`
	Title     *string     `json:"title"`
	Content   string      `json:"content"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	User      UserSummary `json:"user"`
}

type CreateReviewRequest struct {
	Rating  int     `json:"rating"`
	Content string  `json:"content"`
	Title   *string `json:"title"`
}

type UpdateReviewRequest struct {
	Rating  *int    `json:"rating"`
	Content *string `json:"content"`
	Title   *string `json:"title"`
}

type ReviewListResponse struct {
	Data       []ReviewResponse `json:"data"`
	Pagination Pagination       `json:"pagination"`
}

func (r *Review) ToResponse(user *User) ReviewResponse {
	resp := ReviewResponse{
		ID:        r.ID,
		BookID:    r.BookID.String(),
		UserID:    r.UserID.String(),
		Rating:    r.Rating,
		Content:   r.Content,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}

	if r.Title.Valid {
		resp.Title = &r.Title.String
	}

	if user != nil {
		resp.User = user.ToSummary()
	}

	return resp
}
