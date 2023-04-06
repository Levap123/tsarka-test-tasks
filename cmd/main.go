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
	"github.com/Levap123/tsarka-test-tasks/internal/services"
	httpserver "github.com/Levap123/tsarka-test-tasks/pkg/http_server"
)

func main() {
	services := services.NewServices()
	handlers := handlers.NewHandlers(services)

	srv := httpserver.New(handlers.InitRoutes(), ":8080")
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

	log.Println("server started")
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("fatal in starting server: %s", err.Error())
	}

}
