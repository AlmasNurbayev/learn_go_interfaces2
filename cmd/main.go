package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"interfaces2/internal/parser"
	"interfaces2/internal/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sqlx.Connect("postgres", "postgres://ps5:PsX314159_@localhost:5432/green")
	if err != nil {
		log.Printf("[ERROR] failed to connect to db: %v", err)
		return
	}
	defer db.Close()

	userStorage := storage.NewUsersStorage(db)
	articleStorage := storage.NewArticlesStorage(db)
	parser := parser.New(
		userStorage,
		articleStorage,
		time.Duration(1)*time.Minute,
	)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func(ctx context.Context) {
		log.Print("started goroutine")

		if err := parser.Start(ctx); err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Printf("[ERROR] failed to run parser: %v", err)
				return
			}

			log.Printf("[INFO] parser stopped")
		}
	}(ctx)

}
