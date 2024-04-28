package grpc

import (
	"context"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/model"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type MessageService interface {
	GetLastMessages(ctx context.Context, n int) ([]model.Message, error)
}

type MessageHandler struct {
	pb.UnimplementedMessageStorageServer
	svc MessageService
}

func NewMessageHandler(svc MessageService) *MessageHandler {
	return &MessageHandler{svc: svc}
}

func (h *MessageHandler) GetLastMessages(ctx context.Context, req *pb.GetLastMessagesRequest) (*pb.GetLastMessagesResponse, error) {
	msgs, err := h.svc.GetLastMessages(ctx, int(req.Number))
	if err != nil {
		log.Printf("service: %v", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	res := make([]*pb.Message, len(msgs))
	for i, msg := range msgs {
		res[i] = &pb.Message{
			SentTime: timestamppb.New(msg.SentTime),
			Nickname: msg.Nickname,
			Message:  msg.Message,
		}
	}
	return &pb.GetLastMessagesResponse{Messages: res}, nil
}
