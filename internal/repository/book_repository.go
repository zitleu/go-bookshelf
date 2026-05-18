package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bookshelf/monolith/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) Create(ctx context.Context, book *domain.Book) (bool, error) {
	book.ID = uuid.New()

	query := `
		INSERT INTO books(id, title, author, description, isbn, published_year, created_by)
		VALUES(:id, :title, :author, :description, :isbn, :published_year, :created_by)
	`

	_, err := b.db.NamedExecContext(ctx, query, book)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (b *BookRepository) GetByID(ctx context.Context, id string) (*domain.Book, error) {
	query := `
		SELECT b.*,
			AVG(r.rating) AS average_rating,
			COUNT(r.id)   AS reviews_count
		FROM books b
		LEFT JOIN reviews r ON r.book_id = b.id
		WHERE b.id = $1
		GROUP BY b.id
	`

	var book domain.Book
	err := b.db.GetContext(ctx, &book, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}

	return &book, nil
}

func (b *BookRepository) List(ctx context.Context, filter domain.BookFilter) ([]domain.Book, int, error) {
	allowedSortCols := map[string]string{
		"title":      "b.title",
		"author":     "b.author",
		"created_at": "b.created_at",
		"rating":     "average_rating",
	}
	sortCol, ok := allowedSortCols[filter.Sort]
	if !ok {
		sortCol = "b.created_at"
	}

	order := "DESC"
	if filter.Order == "asc" {
		order = "ASC"
	}

	args := []interface{}{}
	argIdx := 1
	whereClause := ""

	if filter.Search != "" {
		pattern := "%" + filter.Search + "%"
		whereClause = fmt.Sprintf("WHERE b.title ILIKE $%d OR b.author ILIKE $%d", argIdx, argIdx+1)
		args = append(args, pattern, pattern)
		argIdx += 2
	}

	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM books b %s`, whereClause)

	var total int
	if err := b.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 20
	}
	page := filter.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	listQuery := fmt.Sprintf(`
		SELECT b.*,
			AVG(r.rating) AS average_rating,
			COUNT(r.id)   AS reviews_count
		FROM books b
		LEFT JOIN reviews r ON r.book_id = b.id
		%s
		GROUP BY b.id
		ORDER BY %s %s
		LIMIT $%d OFFSET $%d
	`, whereClause, sortCol, order, argIdx, argIdx+1)

	args = append(args, limit, offset)

	var books []domain.Book
	if err := b.db.SelectContext(ctx, &books, listQuery, args...); err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

func (b *BookRepository) Update(ctx context.Context, book *domain.Book) (bool, error) {
	query := `
		UPDATE books
		SET title          = :title,
			author         = :author,
			description    = :description,
			isbn           = :isbn,
			published_year = :published_year,
			updated_at     = NOW()
		WHERE id = :id
	`

	_, err := b.db.NamedExecContext(ctx, query, book)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (b *BookRepository) Delete(ctx context.Context, id string) (bool, error) {
	query := `DELETE FROM books WHERE id = $1`

	_, err := b.db.ExecContext(ctx, query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
