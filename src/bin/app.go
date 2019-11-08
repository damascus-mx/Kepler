package app

import (
	"fmt"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/damascus-mx/kepler/src/core"
	"github.com/damascus-mx/kepler/src/routes"
)

// InitApplication Returns a new chi Router pointer
func InitApplication() *chi.Mux {
	// Create a new Chi router
	fmt.Print("Running new REST Microservice")
	app := chi.NewRouter()

	// Set CORS policies
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Use CORS
	app.Use(cors.Handler)

	// Set default middlewares
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	app.Use(middleware.Timeout(60 * time.Second))

	// Set routes
	var user core.IRoute = routes.UserRoutes{}
	user.SetRoutes(app)

	return app
}
