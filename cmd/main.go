package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Levap123/tsarka-test-tasks/internal/handlers"
	httpserver "github.com/Levap123/tsarka-test-tasks/internal/infrastructure/http"
	"github.com/Levap123/tsarka-test-tasks/internal/infrastructure/postgres"
	"github.com/Levap123/tsarka-test-tasks/internal/infrastructure/redis"
	"github.com/Levap123/tsarka-test-tasks/internal/repository"
	"github.com/Levap123/tsarka-test-tasks/internal/services"
)

// @title			TSARKA test tasks API
// @version		1.0
// @description	Тестовое задние на позицию Golang junior в ЦАРКА
// @host			localhost:8080
// @BasePath		/rest
func main() {
	redisClient, err := redis.New()
	if err != nil {
		log.Fatalf("can not establish connection with redis: %v\n", err)
	}

	DB, err := postgres.New()
	if err != nil {
		log.Fatalf("can not establish connection with postgres: %v\n", err)
	}

	repo := repository.NewRepository(redisClient, DB)
	services := services.NewServices(repo)
	handlers := handlers.NewHandlers(services)

	addr := ":8080"
	srv := httpserver.New(handlers.InitRoutes(), addr)
	srv.SetKeepAlivesEnabled(true)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdown
		log.Println("server stopped")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("error in stopping server: %s", err.Error())
		}
	}()

	log.Printf("server started, go to http://localhost%s/swagger", addr)
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("fatal in starting server: %v", err)
	}
}
