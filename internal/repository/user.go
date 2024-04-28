package repository

import (
	"context"
	"errors"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/model"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) AddUser(ctx context.Context, user model.User) error {
	_, err := r.pool.Exec(ctx, `INSERT INTO users (nickname, password_hash) VALUES ($1, $2)`, user.Nickname, user.PasswordHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return model.ErrAlreadyExists
		}
	}
	return err
}

func (r *UserRepository) GetUser(ctx context.Context, nickname string) (model.User, error) {
	row := r.pool.QueryRow(ctx, `SELECT password_hash FROM users WHERE nickname = $1`, nickname)

	var passwordHash []byte
	err := row.Scan(&passwordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, model.ErrNotFound
		}
		return model.User{}, err
	}

	return model.User{Nickname: nickname, PasswordHash: passwordHash}, nil
}
