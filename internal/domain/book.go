package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID       `db:"id"`
	Title         string          `db:"title"`
	Author        string          `db:"author"`
	CreatedBy     string          `db:"created_by"`
	CreatedAt     time.Time       `db:"created_at"`
	UpdatedAt     time.Time       `db:"updated_at"`
	Description   sql.NullString  `db:"description"`
	ISBN          sql.NullString  `db:"isbn"`
	PublishedYear sql.NullInt32   `db:"published_year"`
	AverageRating sql.NullFloat64 `db:"average_rating"`
	ReviewsCount  int             `db:"reviews_count"`
}

type BookResponse struct {
	ID            uuid.UUID    `json:"id"`
	Title         string       `json:"title"`
	Author        string       `json:"author"`
	CreatedBy     string       `json:"created_by"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	Description   *string      `json:"description"`
	ISBN          *string      `json:"isbn"`
	PublishedYear *int32       `json:"published_year"`
	AverageRating *float64     `json:"average_rating"`
	ReviewsCount  int          `json:"reviews_count"`
	Creator       *UserSummary `json:"creator,omitempty"`
}

type CreateBookRequest struct {
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Description   *string `json:"description"`
	ISBN          *string `json:"isbn"`
	PublishedYear *int    `json:"published_year"`
}

type UpdateBookRequest struct {
	Title         *string `json:"title"`
	Author        *string `json:"author"`
	Description   *string `json:"description"`
	ISBN          *string `json:"isbn"`
	PublishedYear *int    `json:"published_year"`
}

type BookFilter struct {
	Search string
	Sort   string
	Order  string
	Page   int
	Limit  int
}

type BookListResponse struct {
	Data       []BookResponse `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

func (b *Book) ToResponse() BookResponse {
	var description *string
	if b.Description.Valid {
		description = &b.Description.String
	}

	var isbn *string
	if b.ISBN.Valid {
		isbn = &b.ISBN.String
	}
	var publishedYear *int32
	if b.PublishedYear.Valid {
		publishedYear = &b.PublishedYear.Int32
	}

	var averageRating *float64
	if b.AverageRating.Valid {
		averageRating = &b.AverageRating.Float64
	}

	return BookResponse{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		CreatedBy:     b.CreatedBy,
		CreatedAt:     b.CreatedAt,
		UpdatedAt:     b.UpdatedAt,
		Description:   description,
		ISBN:          isbn,
		PublishedYear: publishedYear,
		AverageRating: averageRating,
		ReviewsCount:  b.ReviewsCount,
	}
}
