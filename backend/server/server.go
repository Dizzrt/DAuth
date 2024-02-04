package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	v1 "github.com/Dizzrt/DAuth/backend/api/v1"
	"github.com/Dizzrt/DAuth/backend/common/dlog"
	"github.com/Dizzrt/DAuth/backend/server/routers"
	v1pb "github.com/Dizzrt/DAuth/proto/generated/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Server struct {
	Gin        *gin.Engine
	GrpcServer *grpc.Server

	s  *http.Server
	wg sync.WaitGroup
}

func NewServer(ctx context.Context) (*Server, error) {
	server := &Server{
		Gin: routers.GetRouter(),
	}

	// init grpc server
	server.GrpcServer = grpc.NewServer(
		grpc.MaxRecvMsgSize(100*1024*1024),
		grpc.InitialConnWindowSize(100000000),
		grpc.InitialConnWindowSize(100000000),
		// grpc.ChainUnaryInterceptor()
		// grpc.ChainStreamInterceptor()
	)

	// register grpc services
	v1pb.RegisterPingServiceServer(server.GrpcServer, &v1.PingService{})

	return server, nil
}

func (server *Server) Run(ctx context.Context, port uint16) error {
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: server.Gin,
	}
	server.s = s

	server.wg.Add(1)
	go func() {
		defer server.wg.Done()

		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			dlog.Fatalf("start server failed with error: %v", err)
		}
	}()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port+1))
	if err != nil {
		return err
	}

	go func() {
		if err := server.GrpcServer.Serve(listen); err != nil {
			dlog.Error("grpc server listen error", err)
		}
	}()

	return nil
}

func (server *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := server.s.Shutdown(ctx)
	if err != nil {
		return err
	}

	server.wg.Wait()

	return nil
}
