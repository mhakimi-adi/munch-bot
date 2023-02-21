package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mhakimi-adi/munch-bot/pkg/config"
)

// @title           Munch-Bot API
// @version         0.1
// @description     Define recipes and generate grocery lists.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()
	config.App(router)
	router.Run("localhost:8080")
}
