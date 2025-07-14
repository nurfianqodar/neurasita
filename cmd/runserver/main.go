package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nurfianqodar/neurasita/internal/handler"
	"github.com/nurfianqodar/neurasita/internal/repository"
	"github.com/nurfianqodar/neurasita/internal/service"
	"github.com/nurfianqodar/neurasita/pkg/config"
)

func main() {
	// Setup database connection
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tx, err := pgxpool.New(ctx, config.DATABASE_URL)
	if err != nil {
		log.Fatalf("[fatal] database connection error: %v\n", err)
	}
	q := repository.New(tx)

	// Setup apps
	userService := service.NewUserService(q)
	userHandler := handler.NewUserHandler(userService)

	// Create and register router
	mux := http.NewServeMux()
	userHandler.RegisterRouter(mux)

	// Implement global middleware here
	var h http.Handler = mux
	// ... another middleware wrap h with middleware

	// Create server
	addr := fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	log.Printf("[info] server listening on http://%s\n", addr)
	s := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	// handle graceful shutdown
	go handleShutdown(s)

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("[fatal] running server error: %v\n", err)
	}

}

func handleShutdown(s *http.Server) {
	// TODO handle graceful shutdown
}
