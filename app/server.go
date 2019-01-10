package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	return &Server{
		router: router,
	}
}

func (s *Server) Run(port int) error {
	s.routes()
	addr := fmt.Sprintf(":%d", port)
	return s.router.Run(addr)
}
