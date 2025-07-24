package main

import (
	"fmt"
	configuration "go_fiber_core_project_api/configuration/app"
)

func main() {

	// Initial configuration
	app_configuration := configuration.NewConfiguration()

	// Initialize router
	app := router.New(db_pool)

	app.Listen(fmt.Sprintf("%s:%d", app_configuration.AppHost, app_configuration.AppPort))
}
