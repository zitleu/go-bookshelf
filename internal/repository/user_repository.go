package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bookshelf/monolith/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	user.ID = uuid.New()

	query := `
		INSERT INTO users(id, username, email, password_hash,created_at)
		VALUES(:id, :username, :email, :password_hash, :created_at)
	`

	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := `
			SELECT *
			FROM users
			WHERE id = $1
		`

	var user domain.User

	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil

}
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
			SELECT *
			FROM users
			WHERE email = $1
		`

	var user domain.User

	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil

}
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := `
			SELECT *
			FROM users
			WHERE username = $1
		`

	var user domain.User

	err := r.db.GetContext(ctx, &user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}
func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users
		SET username = :username,
			email = :email,
			password_hash = :password_hash
		WHERE id = :id
	`

	if _, err := r.db.NamedExecContext(ctx, query, user); err != nil {
		return err
	}

	return nil

}
func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	query := `
		SELECT exists(
			SELECT 1
			FROM users
			where email = $1
		)
	`

	var exists bool
	err := r.db.GetContext(ctx, &exists, query, email)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func (r *UserRepository) UsernameExists(ctx context.Context, username string) (bool, error) {
	query := `
		SELECT exists(
			SELECT 1
			FROM users
			where username = $1
		)
	`

	var exists bool
	err := r.db.GetContext(ctx, &exists, query, username)
	if err != nil {
		return false, err
	}

	return exists, nil
}
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
