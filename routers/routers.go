package routers

import (
	"github.com/gofiber/contrib/fiberi18n"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"golang.org/x/text/language"
)

func New(db_pool *sqlx.DB) *fiber.App {
	f := fiber.New(fiber.Config{})

	f.Use(logger.New())

	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	})).Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath: "pkg/translates/localize/i18n",
			AcceptLanguages: []language.Tag{
				language.Chinese,
				language.MustParse("km"),
				language.English,
			},
			DefaultLanguage: language.Khmer,
		}),
	)
	return f
}
