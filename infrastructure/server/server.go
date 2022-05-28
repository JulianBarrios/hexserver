package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	health "github.com/julianbarrios/hexserver/infrastructure/server/handlers"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("")
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler)
}
