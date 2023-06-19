package main

import (
	"backend/cmd"
	"backend/internal/logger"

	_ "backend/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host petstore.swagger.io
// @BasePath /v2
func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		logger.Fatal("Failed to Start")
	}
}
