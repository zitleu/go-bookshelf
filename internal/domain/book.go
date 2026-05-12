package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID       `json:"id" db:"id"`
	Title         string          `json:"title" db:"title"`
	Author        string          `json:"author" db:"author"`
	CreatedBy     string          `json:"created_by" db:"created_by"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at" db:"updated_at"`
	Description   sql.NullString  `json:"description" db:"description"`
	ISBN          sql.NullString  `json:"isbn" db:"isbn"`
	PublishedYear sql.NullInt32   `json:"published_year" db:"published_year"`
	AverageRating sql.NullFloat64 `json:"average_rating" db:"average_rating"`
	ReviewsCount  int             `json:"reviews_count" db:"reviews_count"`
}

type BookResponse struct {
	ID            uuid.UUID   `json:"id" db:"id"`
	Title         string      `json:"title" db:"title"`
	Author        string      `json:"author" db:"author"`
	CreatedBy     string      `json:"created_by" db:"created_by"`
	CreatedAt     time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time  `json:"updated_at" db:"updated_at"`
	Description   *string     `json:"description" db:"description"`
	ISBN          *string     `json:"isbn" db:"isbn"`
	PublishedYear *int        `json:"published_year" db:"published_year"`
	AverageRating *float64    `json:"average_rating" db:"average_rating"`
	ReviewsCount  int         `json:"reviews_count" db:"reviews_count"`
	Creator       UserSummary `json:"creator,omitempty"`
}

type CreateBookRequest struct {
	Title         string
	Author        string
	Description   *string
	ISBN          *string
	PublishedYear *int
}

type UpdateBookRequest struct {
	ID            *uuid.UUID      `json:"id" db:"id"`
	Title         *string         `json:"title" db:"title"`
	Author        *string         `json:"author" db:"author"`
	CreatedBy     *string         `json:"created_by" db:"created_by"`
	CreatedAt     *time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at" db:"updated_at"`
	Description   *sql.NullString `json:"description" db:"description"`
	ISBN          *sql.NullString `json:"isbn" db:"isbn"`
	PublishedYear *sql.NullInt32  `json:"published_year" db:"published_year"`
}

type BookFilter struct {
	Search []string
	Sort   []string
	Order  []string
	Page   []string
	Limit  []string
}

type BookListResponse struct {
	Data []BookResponse
	Pagination Pagination
}
