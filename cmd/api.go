package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/suliman_dahan/go-ecom-api/internal/adapters/postgresql/sqlc"
	"github.com/suliman_dahan/go-ecom-api/internal/products"
)

// mount
func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	// middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(middleware.Timeout(time.Second * 60))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	router.Get("/products", productHandler.ListProducts)

	// http.ListenAndServe(app.Config.address, router)
	return router
}

// run

func (app *application) run(h http.Handler) error {
	server := &http.Server{
		Addr:         app.Config.address,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	slog.Info("Server has been started on port 8080")
	return server.ListenAndServe()
}

type application struct {
	Config config
	db     *pgx.Conn
}

type config struct {
	address string
	db      dbConfig
}

type dbConfig struct {
	dsn string // dsn -> domain string
}
