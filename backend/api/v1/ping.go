package v1

import (
	"context"

	v1 "github.com/Dizzrt/DAuth/proto/generated/v1"
)

type PingService struct {
	v1.UnimplementedPingServiceServer
}

func (s *PingService) Ping(ctx context.Context, in *v1.PingRequest) (*v1.PingResponse, error) {
	return &v1.PingResponse{
		Pong: "pong",
	}, nil
}
