package server

import "github.com/gin-gonic/gin"

type Server struct {
	Gin *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		Gin: gin.Default(),
	}

	return server
}

func (server *Server) Run() {

}
