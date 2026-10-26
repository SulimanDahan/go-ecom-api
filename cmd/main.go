package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/suliman_dahan/go-ecom-api/internal/env"
)

func main() {
	ctx := context.Background()

	cfg := config{
		address: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=1234 dbname=ecom sslmode=disable"),
		},
	}

	// Structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
		// os.Exit(1)
	}

	logger.Info("Connected to database", "dsn", cfg.db.dsn)

	defer conn.Close(ctx)

	api := application{
		Config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
