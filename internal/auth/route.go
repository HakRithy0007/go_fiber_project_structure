package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type AuthRoute struct {
	app     *fiber.App
	db_pool *sqlx.DB
	handler *AuthHandler
}

func NewRoute(app *fiber.App, db_pool *sqlx.DB) *AuthRoute {
	handler := NewHandler(db_pool)
	return &AuthRoute{
		app:     app,
		db_pool: db_pool,
		handler: handler,
	}
}

func (a *AuthRoute) RegisterAuthRoute() *AuthRoute {
	v1 := a.app.Group("/api/v1")
	auth := v1.Group("/auth")
	auth.Post("/login", a.handler.Login)

	return a
}
