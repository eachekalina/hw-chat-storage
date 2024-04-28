package service

import (
	"context"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/model"
)

type UserRepository interface {
	AddUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, nickname string) (model.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser(ctx context.Context, user model.User) error {
	return s.repo.AddUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, nickname string) (model.User, error) {
	return s.repo.GetUser(ctx, nickname)
}
