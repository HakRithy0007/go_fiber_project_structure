package handler

import (
	"go_fiber_core_project_api/internal/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type ServiceHandlers struct {
	Fronted *FrontService
}

type FrontService struct {
	AuthHandler *auth.AuthRoute
}

func NewFrontService(app *fiber.App, db_pool *sqlx.DB) *FrontService {

	// Authentication
	auth := auth.NewRoute(app, db_pool).RegisterAuthRoute()

	return &FrontService{
		AuthHandler: auth,
	}
}

func NewServiceHandlers(app *fiber.App, db_pool *sqlx.DB) *ServiceHandlers {

	return &ServiceHandlers{
		Fronted: NewFrontService(app, db_pool),
	}
}
