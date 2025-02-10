package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1routes "github.com/gsaudx/api-go-project/internal/api/v1/routes"
	"github.com/gsaudx/api-go-project/internal/config"
	"github.com/gsaudx/api-go-project/internal/middlewares"
	pkgmiddlewares "github.com/gsaudx/api-go-project/pkg/middlewares"
)

func main() {
	// Load configuration from environment.
	cfg := config.Load()

	// Create a new HTTP router (ServeMux).
	mux := http.NewServeMux()

	// Use the aggregated routes loader for v1 endpoints.
	v1routes.RegisterRoutes(mux)

	// Compose middleware: wrap your mux with logging, request tracing, and authentication.
	// The order is important because it determines the flow of requests.
	handler := pkgmiddlewares.LoggingMiddleware(
		pkgmiddlewares.RequestTracingMiddleware(
			middlewares.AuthMiddleware(mux),
		),
	)

	// Create the HTTP server with proper timeouts.
	server := &http.Server{
		Addr:         cfg.Address(),
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server in a separate goroutine.
	go func() {
		log.Printf("Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen error: %v", err)
		}
	}()

	// Wait for an interrupt to gracefully shut down.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Attempt graceful shutdown within 15 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}
	log.Println("Server stopped gracefully")
}
