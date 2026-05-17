package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bookshelf/monolith/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func Create(ctx context.Context) {}
func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := `
			SELECT *
			FROM users
			WHERE id = $1
		`

	var user domain.User

	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil

}
func GetByEmail(ctx context.Context)     {}
func GetByUsername(ctx context.Context)  {}
func Update(ctx context.Context)         {}
func EmailExists(ctx context.Context)    {}
func UsernameExists(ctx context.Context) {}
func (r *UserRepository) GetByIDs(ctx context.Context, ids []string) (map[string]*domain.User, error) {
	query, args, err := sqlx.In("SELECT * FROM users WHERE id in (?)", ids)
	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)

	var users []domain.User
	if err := r.db.SelectContext(ctx, &users, query, args...); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	result := make(map[string]*domain.User, len(users))
	for _, v := range users {
		result[v.ID.String()] = &v
	}

	return result, nil
}
