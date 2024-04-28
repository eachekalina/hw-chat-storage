package repository

import (
	"context"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessageRepository struct {
	pool *pgxpool.Pool
}

func NewMessageRepository(pool *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{pool: pool}
}

func (r *MessageRepository) AddMessage(ctx context.Context, msg model.Message) error {
	_, err := r.pool.Exec(ctx, `INSERT INTO messages (sent_time, nickname, msg) VALUES ($1, $2, $3)`, msg.SentTime, msg.Nickname, msg.Message)
	return err
}

func (r *MessageRepository) GetLastMessages(ctx context.Context, n int) ([]model.Message, error) {
	rows, err := r.pool.Query(ctx, `SELECT sent_time, nickname, message FROM messages ORDER BY sent_time DESC LIMIT $1`, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Message])
}
