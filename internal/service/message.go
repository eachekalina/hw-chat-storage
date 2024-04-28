package service

import (
	"context"
	"fmt"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/model"
	"log"
)

type MessageRepository interface {
	AddMessage(ctx context.Context, msg model.Message) error
	GetLastMessages(ctx context.Context, n int) ([]model.Message, error)
}

type MessageCache interface {
	AddMessage(ctx context.Context, msg model.Message) error
	GetLastMessages(ctx context.Context, n int) ([]model.Message, error)
}

type MessageService struct {
	repo  MessageRepository
	cache MessageCache
}

func NewMessageService(repo MessageRepository, cache MessageCache) *MessageService {
	return &MessageService{
		repo:  repo,
		cache: cache,
	}
}

func (s *MessageService) AddMessage(ctx context.Context, msg model.Message) error {
	err := s.repo.AddMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("message repo: %w", err)
	}
	err = s.cache.AddMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("message cache: %w", err)
	}
	return nil
}

func (s *MessageService) GetLastMessages(ctx context.Context, n int) ([]model.Message, error) {
	msgs, err := s.cache.GetLastMessages(ctx, n)
	if err == nil {
		return msgs, nil
	}
	log.Printf("message cache: %v", err)
	msgs, err = s.repo.GetLastMessages(ctx, n)
	if err != nil {
		return nil, fmt.Errorf("message repo: %w", err)
	}
	return msgs, nil
}
