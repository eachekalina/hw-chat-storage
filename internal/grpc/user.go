package grpc

import (
	"context"
	"errors"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/model"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type UserService interface {
	AddUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, nickname string) (model.User, error)
}

type UserHandler struct {
	pb.UnimplementedUserStorageServer
	svc UserService
}

func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	user := model.User{
		Nickname:     req.User.Nickname,
		PasswordHash: req.User.PasswordHash,
	}

	err := h.svc.AddUser(ctx, user)
	if err != nil {
		if errors.Is(err, model.ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "already exists")
		}
		log.Printf("service: %v", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pb.AddUserResponse{}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.svc.GetUser(ctx, req.Nickname)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "not found")
		}
		log.Printf("service: %v", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	res := &pb.GetUserResponse{User: &pb.User{
		Nickname:     user.Nickname,
		PasswordHash: user.PasswordHash,
	}}
	return res, nil
}
