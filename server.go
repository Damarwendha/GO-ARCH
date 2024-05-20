package main

import (
	"database/sql"
	"fmt"
	"go-arch/config"
	"go-arch/controller"
	"go-arch/middleware"
	"go-arch/repository"
	"log"

	"go-arch/service"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

// Note: All service should be injected
type Server struct {
	jwtService    service.JwtServiceI
	authService   service.AuthServiceI
	authorService service.AuthorServiceI
	taskService   service.TaskServiceI
	engine        *gin.Engine
}

func (s *Server) initRoute() {
	router := s.engine.Group("/api/v1")
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)

	controller.NewAuthController(s.authService, router, authMiddleware).Routing()
	controller.NewAuthorController(s.authorService, router, authMiddleware).Routing()
	controller.NewTaskController(s.taskService, router, authMiddleware).Routing()
}

func (s *Server) Run() {
	s.initRoute()

	s.engine.Run(":8080")
}

func NewServer() *Server {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("Unexpected error di NewServer", err)
	}

	urlConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.DbUser, c.DbPassword, c.DbName, c.DbPort)

	database, err := sql.Open(c.Driver, urlConnection)
	if err != nil {
		log.Fatal(err)
	}

	authorRepo := repository.NewAuthorRepo(database)
	authorService := service.NewAuthorService(authorRepo)

	jwtService := service.NewJwtService(c.TokenConfig)
	authService := service.NewAuthService(jwtService, authorService)

	taskRepo := repository.NewTaskRepo(database)
	taskService := service.NewTaskService(taskRepo)

	// Note: All service should be injected
	return &Server{
		authService:   authService,
		authorService: authorService,
		engine:        gin.Default(),
		taskService:   taskService,
		jwtService:    jwtService,
	}
}
