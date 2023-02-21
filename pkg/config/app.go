package config

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mhakimi-adi/munch-bot/docs/swagger"
	"github.com/mhakimi-adi/munch-bot/pkg/controller"
	"github.com/mhakimi-adi/munch-bot/pkg/route"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func App(router *gin.Engine) {
	c := controller.NewController()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler, url),
	)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			route.Item(v1.Group("/item"), c)
		}
	}
}
