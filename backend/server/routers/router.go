package routers

import (
	"context"
	"net/http"
	"time"

	v1 "github.com/Dizzrt/DAuth/proto/generated/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer conn.Close()
		c := v1.NewPingServiceClient(conn)

		cctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.Ping(cctx, &v1.PingRequest{})
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": r.GetPong(),
		})
	})

	// load routers here

	return r
}
