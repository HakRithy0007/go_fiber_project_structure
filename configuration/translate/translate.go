package translate

import (
	"fmt"
	custom_error "go_fiber_core_project_api/pkg/utils/errors"
	loggers "go_fiber_core_project_api/pkg/utils/loggers"
	"log"
	"path/filepath"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var bundle *i18n.Bundle

func InitTranslate() *custom_error.ErrorResponse {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	localeFiles := []string{
		"pkg/translates/localize/i18n/en.yaml",
		"pkg/translates/localize/i18n/km.yaml",
	}

	for _, file := range localeFiles {
		_, err := bundle.LoadMessageFile(filepath.Join(file))
		if err != nil {
			log.Printf("Error loading local file %s: %v", file, err)
			loggers.NewCustomLog("translate_error", err.Error(), "error")
			return &custom_error.ErrorResponse{
				MessageID: "errorLoadMessage",
				Err:       err,
			}
		}
	}
	return nil
}

func TranslateWithError(c *fiber.Ctx, key string, templateData ...map[string]interface{}) (string, *custom_error.ErrorResponse) {
	if bundle == nil {
		loggers.NewCustomLog("I18nNotInit", InitTranslate().ErrorString(), "error")
		return "", &custom_error.ErrorResponse{
			MessageID: key,
			Err:       fmt.Errorf("translation service is unavailable for MessageID: %s", key),
		}
	}

	lang := c.Get("Accept-Language", "en")
	localizer := i18n.NewLocalizer(bundle, lang)

	data := map[string]interface{}{}
	if len(templateData) > 0 && templateData[0] != nil {
		data = templateData[0]
	}

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: data,
	})
	if err != nil {
		log.Printf("Error localizing message ID %s: %v", key, err)
		loggers.NewCustomLog("TranslationNotFound", err.Error(), "error")
		return "", &custom_error.ErrorResponse{
			MessageID: key,
			Err:       fmt.Errorf("translation not found for MessageID: %s", key),
		}
	}
	return msg, nil
}

func Translate(c *fiber.Ctx, key string) string {
	return fiberi18n.MustLocalize(c, &i18n.LocalizeConfig{
		MessageID: key,
	})
}
