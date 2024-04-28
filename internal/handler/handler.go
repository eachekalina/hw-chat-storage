package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/model"
)

type Service interface {
	AddMessage(ctx context.Context, msg model.Message) error
}

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h Handler) HandleMessage(ctx context.Context, content []byte) error {
	msg, err := parseMessage(content)
	if err != nil {
		return fmt.Errorf("parse message: %w", err)
	}
	return h.svc.AddMessage(ctx, msg)
}

func parseMessage(content []byte) (model.Message, error) {
	var msg model.Message
	err := json.Unmarshal(content, &msg)
	return msg, err
}
