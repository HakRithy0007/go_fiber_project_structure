package auth

import (
	custom_translate "go_fiber_core_project_api/configuration/translate"
	constants "go_fiber_core_project_api/pkg/constants"
	responses "go_fiber_core_project_api/pkg/utils/responses"
	custom_validator "go_fiber_core_project_api/pkg/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type AuthHandler struct {
	db_pool     *sqlx.DB
	authService AuthService
}

func NewHandler(db_pool *sqlx.DB) *AuthHandler {
	return &AuthHandler{
		db_pool:     db_pool,
		authService: NewService(db_pool),
	}
}

func (a *AuthHandler) Login(c *fiber.Ctx) error {
	browser := c.Get("User-Agent", "unknown")
	clientIP := c.IP()
	v := custom_validator.NewValidator()
	req := &AuthLoginRequest{}

	if err := req.bind(c, v); err != nil {
		msg, err_msg := custom_translate.TranslateWithError(c, "login_invalid")
		if err_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				responses.NewResponseError(
					err_msg.ErrorString(),
					constants.Translate_failed,
					err_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			responses.NewResponseError(
				msg,
				constants.Login_invalid,
				err,
			),
		)
	}

	success, err := a.authService.Login(req.Auth.Username, req.Auth.Password, browser, clientIP)
	if err != nil {
		msg, msg_err := custom_translate.TranslateWithError(c, err.MessageID)
		if msg_err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.NewResponseError(
				msg_err.Err.Error(),
				constants.Translate_failed,
				msg_err.Err,
			))
		}
		return c.Status(fiber.StatusUnauthorized).JSON(responses.NewResponseError(
			msg,
			constants.Login_failed,
			err.Err,
		))
	} else {
		msg, err_msg := custom_translate.TranslateWithError(c, "login_success")
		if err_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(responses.NewResponseError(
				err_msg.ErrorString(),
				constants.Translate_failed,
				err_msg.Err,
			))
		}
		return c.Status(fiber.StatusOK).JSON(responses.NewResponse(
			msg,
			constants.Login_success,
			success,
		))
	}
}
