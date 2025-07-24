package auth

import (
	custom_error "go_fiber_core_project_api/pkg/utils/errors"

	"github.com/jmoiron/sqlx"
)

type AuthService interface {
	CheckSession(loginSession string, userID float64) (bool, *custom_error.ErrorResponse)
	Login(username, password, browser, clientIP string) (*AuthResponse, *custom_error.ErrorResponse)
}

type ServiceImpl struct {
	db_pool *sqlx.DB
	repo    AuthRepo
}

func NewService(db_pool *sqlx.DB) AuthService {
	repo := NewRepository(db_pool)
	return &ServiceImpl{
		db_pool: db_pool,
		repo:    repo,
	}
}

func (a *ServiceImpl) Login(username, password, browser, clientIP string) (*AuthResponse, *custom_error.ErrorResponse) {
	response, err := a.repo.Login(username, password, browser, clientIP)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (a *ServiceImpl) CheckSession(loginSession string, userID float64) (bool, *custom_error.ErrorResponse) {
	return a.repo.CheckSession(loginSession, userID)
}
