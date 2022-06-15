package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	mooc "github.com/julianbarrios/hexserver/infrastructure/data"
	"github.com/julianbarrios/hexserver/infrastructure/server/handlers/course"
	"github.com/julianbarrios/hexserver/infrastructure/server/handlers/health"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//dependencies
	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		engine:           gin.New(),
		httpAddr:         fmt.Sprintf("%s:%d", host, port),
		courseRepository: courseRepository,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("")
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/course", course.CreateCourse(s.courseRepository))
}
