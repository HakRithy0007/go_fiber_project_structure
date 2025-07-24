package auth

import (
	custom_validator "go_fiber_core_project_api/pkg/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
)

// login
type AuthLoginRequest struct {
	Auth struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"auth"`
}

type AuthResponse struct {
	Auth struct {
		Token     string `json:"token"`
		TokenType string `json:"token_type"`
	} `json:"auths"`
}

type MemberData struct {
	ID       int    `db:"id"`
	Username string `db:"user_name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type RedisSession struct {
	LoginSession string `json:"login_session"`
}

type MiniLogin struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	SecretKey string `json:"secret_key"`
}

type LoginMiniResponse struct {
	Success    bool              `json:"success"`
	Message    string            `json:"message"`
	StatusCode int               `json:"status_code"`
	Data       LoginResponseBody `json:"data"`
}
type LoginResponseBody struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

type LoginMiniRes struct {
	Id             float64 `json:"id"`
	Username       string  `json:"user_name"`
	MembershipId   float64 `json:"membership_id"`
	MembershipRole string  `json:"membership_role"`
	RoleId         float64 `json:"role_id"`
}

// ! test with Mini
type UserloginMini struct {
	Username  string `json:"user_name"`
	Password  string `json:"password"`
	SecretKey string `json:"secret_key"`
}

type MiniResponse struct {
	Id           float64 `json:"id"`
	Username     string  `json:"user_name"`
	MembershipId float64 `json:"membership_id"`
	RoleId       float64 `json:"role_id"`
}

type MiniLoginResponse struct {
	Success    bool              `json:"success"`
	Message    string            `json:"message"`
	StatusCode int               `json:"status_code"`
	Data       LoginResponseBody `json:"data"`
}

type MemberLoginResponse struct {
	Token string `json:"token"`
}

type MemberSyncReq struct {
	ID             int             `json:"id"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Username       string          `json:"user_name"`
	Password       string          `json:"password"`
	Email          string          `json:"email"`
	LoginSession   *string         `json:"login_session"`
	ProfilePhoto   *string         `json:"profile_photo"`
	MemberAlias    *string         `json:"member_alias"`
	PhoneNumber    *string         `json:"phone_number"`
	MemberAvatarID *int            `json:"member_avatar_id"`
	Commission     decimal.Decimal `json:"commission"`
	StatusID       int             `json:"status_id"`
	Order          int             `json:"order"`
	CreatedBy      int             `json:"created_by"`
	CreateAt       string          `json:"created_at"`
}

func (r *AuthLoginRequest) bind(c *fiber.Ctx, v *custom_validator.Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}
