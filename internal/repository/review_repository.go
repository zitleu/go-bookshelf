package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bookshelf/monolith/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ReviewRepository struct {
	db *sqlx.DB
}

func NewReviewRepository(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(ctx context.Context, review *domain.Review) (bool, error) {
	review.ID = uuid.New()

	query := `
		INSERT INTO reviews(id, book_id, user_id, rating, title, content)
		VALUES(:id, :book_id, :user_id, :rating, :title, :content)
	`

	_, err := r.db.NamedExecContext(ctx, query, review)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ReviewRepository) GetByID(ctx context.Context, id string) (*domain.Review, error) {
	query := `SELECT * FROM reviews WHERE id = $1`

	var review domain.Review
	err := r.db.GetContext(ctx, &review, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrReviewNotFound
		}
		return nil, err
	}

	return &review, nil
}

func (r *ReviewRepository) ListByBookID(ctx context.Context, bookID string, page, limit int) ([]domain.Review, int, error) {
	if limit <= 0 {
		limit = 20
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	var total int
	if err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM reviews WHERE book_id = $1`, bookID).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT * FROM reviews
		WHERE book_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	var reviews []domain.Review
	if err := r.db.SelectContext(ctx, &reviews, query, bookID, limit, offset); err != nil {
		return nil, 0, err
	}

	return reviews, total, nil
}

func (r *ReviewRepository) Update(ctx context.Context, review *domain.Review) (bool, error) {
	query := `
		UPDATE reviews
		SET rating     = :rating,
			title      = :title,
			content    = :content,
			updated_at = NOW()
		WHERE id = :id
	`

	_, err := r.db.NamedExecContext(ctx, query, review)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ReviewRepository) Delete(ctx context.Context, id string) (bool, error) {
	query := `DELETE FROM reviews WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ReviewRepository) UserHasReviewedBook(ctx context.Context, userID, bookID string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM reviews
			WHERE user_id = $1 AND book_id = $2
		)
	`

	var exists bool
	if err := r.db.QueryRowContext(ctx, query, userID, bookID).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
