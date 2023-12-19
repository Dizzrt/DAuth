package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Dizzrt/DAuth/backend/common/dlog"
	"github.com/Dizzrt/DAuth/backend/server/routers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *gin.Engine

	s  *http.Server
	wg sync.WaitGroup
}

func NewServer(ctx context.Context) (*Server, error) {
	server := &Server{
		Gin: routers.GetRouter(),
	}

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
