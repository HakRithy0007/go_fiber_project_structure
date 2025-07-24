package auth

import (
	"database/sql"
	"errors"
	"fmt"
	app "go_fiber_core_project_api/configuration/app"
	custom_error "go_fiber_core_project_api/pkg/utils/errors"
	custom_logger "go_fiber_core_project_api/pkg/utils/loggers"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type AuthRepo interface {
	Login(username, password string, userAgent string, clientIP string) (*AuthResponse, *custom_error.ErrorResponse)
	CheckSession(login_session string, userID float64) (bool, *custom_error.ErrorResponse)
}
type AuthRepoImpl struct {
	db_pool *sqlx.DB
}

func NewRepository(db_pool *sqlx.DB) *AuthRepoImpl {
	return &AuthRepoImpl{
		db_pool: db_pool,
	}
}

func (a *AuthRepoImpl) Login(username, password string, userAgent string, clientIP string) (*AuthResponse, *custom_error.ErrorResponse) {
	var member MemberData
	msg := custom_error.ErrorResponse{}

	query :=
		`
			SELECT id,  user_name, email, password
			FROM tbl_members 
			WHERE user_name = $1 AND password = $2 AND deleted_at is NULL
		`

	err := a.db_pool.Get(&member, query, username, password)
	if err != nil {
		custom_logger.NewCustomLog("member_not_found", err.Error(), "error")
		return nil, msg.NewErrorResponse("member_not_found", fmt.Errorf("member not found. Please check the provided information"))
	}

	var res AuthResponse

	hours := app.GetenvInt("JWT_EXP_HOUR", 7)
	expirationTime := time.Now().Add(time.Duration(hours) * time.Hour)
	login_session, err := uuid.NewV7()

	if err != nil {
		custom_logger.NewCustomLog("uuid_generate_failed", err.Error(), "error")
		return nil, msg.NewErrorResponse("uuid_generate_failed", fmt.Errorf("failed to generate UUID. Please try again later"))
	}

	claims := jwt.MapClaims{
		"player_id":     member.ID,
		"username":      member.Username,
		"login_session": login_session.String(),
		"exp":           expirationTime.Unix(),
	}

	errs := godotenv.Load()
	if errs != nil {
		custom_logger.NewCustomLog("error_load_env", errs.Error(), "error")
	}

	secret_key := os.Getenv("JWT_SECRET_KEY")

	updateQuery :=
		`UPDATE  tbl_memberSET login_session = $WHERE id = $2`
	_, err = a.db_pool.Exec(updateQuery, login_session.String(), member.ID)
	if err != nil {
		custom_logger.NewCustomLog("session_update_failed", err.Error(), "error")
		return nil, msg.NewErrorResponse("session_update_failed", fmt.Errorf("cannot update session"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		custom_logger.NewCustomLog("jwt_failed", err.Error(), "error")
		return nil, msg.NewErrorResponse("jwt_failed", fmt.Errorf("failed to get jwt"))
	}

	res.Auth.Token = tokenString
	res.Auth.TokenType = "jwt"

	return &res, nil
}

func (a *AuthRepoImpl) CheckSession(login_session string, memberID float64) (bool, *custom_error.ErrorResponse) {
	msg := custom_error.ErrorResponse{}
	var loginSession string
	query :=
		`SELECT	login_sessionFROM tbl_membersWHERE login_session = $1`

	err := a.db_pool.Get(&loginSession, query, login_session)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			custom_logger.NewCustomLog("invalid_sesion_id", "invalid login session: "+login_session, "warn")
			return false, msg.NewErrorResponse("invalid_sesion_id", fmt.Errorf("invalid login session"))
		}
		custom_logger.NewCustomLog("query_data_failed", err.Error(), "error")
		return false, msg.NewErrorResponse("query_data_failed", fmt.Errorf("database query error"))
	}

	if loginSession != login_session {
		return false, msg.NewErrorResponse("invalid_sesion_id", fmt.Errorf("invalid login session"))
	}
	return true, nil
}
