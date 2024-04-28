package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/model"
	"github.com/redis/go-redis/v9"
)

const key = "messages"

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{client: client}
}

func (r *Cache) AddMessage(ctx context.Context, msg model.Message) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("json: %w", err)
	}
	err = r.client.LPush(ctx, key, bytes).Err()
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}
	return nil
}

func (r *Cache) GetLastMessages(ctx context.Context, n int) ([]model.Message, error) {
	result, err := r.client.LRange(ctx, key, 0, int64(n-1)).Result()
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}
	if len(result) < n {
		return nil, errors.New("not enough elements")
	}
	msgs := make([]model.Message, len(result))
	for i, data := range result {
		err := json.Unmarshal([]byte(data), &(msgs[i]))
		if err != nil {
			return nil, fmt.Errorf("json: %w", err)
		}
	}
	return msgs, nil
}
