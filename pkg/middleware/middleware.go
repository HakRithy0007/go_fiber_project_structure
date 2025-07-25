package middleware

import (
	"fmt"
	custom_translate "go_fiber_core_project_api/configuration/translate"
	"go_fiber_core_project_api/internal/auth"
	custom_models "go_fiber_core_project_api/pkg/model"
	responses "go_fiber_core_project_api/pkg/utils/responses"
	"log"
	"net/http"
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func NewJwtMiddleWare(app *fiber.App, db_pool *sqlx.DB, redis *redis.Client) {
	errs := godotenv.Load()
	if errs != nil {
		log.Fatalf("Error loading .env file")
	}
	secret_key := os.Getenv("JWT_SECRET_KEY")

	// JWT Middleware (non-websocket only)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret_key)},
		ContextKey: "jwt_data",
	}))

	app.Use(func(c *fiber.Ctx) error {
		user_token := c.Locals("jwt_data").(*jwt.Token)
		uclaim := user_token.Claims.(jwt.MapClaims)

		return handleUserContext(c, uclaim, db_pool, redis)
	})
}

func handleUserContext(c *fiber.Ctx, uclaim jwt.MapClaims, db *sqlx.DB, redis *redis.Client) error {

	login_session, ok := uclaim["login_session"].(string)
	if !ok || login_session == "" {
		smg_error := responses.NewResponseError(
			custom_translate.Translate(c, "login_session_missing"),
			-500,
			fmt.Errorf("%s", custom_translate.Translate(c, "login_session_missing")),
		)
		return c.Status(http.StatusUnprocessableEntity).JSON(smg_error)
	}

	uCtx := custom_models.PlayerContext{
		PlayerID:     uclaim["player_id"].(float64),
		UserName:     uclaim["username"].(string),
		LoginSession: uclaim["login_session"].(string),
		Exp:          time.Unix(int64(uclaim["exp"].(float64)), 0),
		UserAgent:    c.Get("User-Agent", "unknown"),
		Ip:           c.IP(),
	}
	c.Locals("PlayerContext", uCtx)

	sv := auth.NewService(db, redis)
	success, err := sv.CheckSession(login_session, uCtx.PlayerID)
	if err != nil || !success {
		smg_error := responses.NewResponseError(
			custom_translate.Translate(c, "login_session_invalid"),
			-500,
			fmt.Errorf("%s", custom_translate.Translate(c, "login_session_invalid")),
		)
		return c.Status(http.StatusUnprocessableEntity).JSON(smg_error)
	}

	return c.Next()
}
