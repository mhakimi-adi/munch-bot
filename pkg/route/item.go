package route

import (
	"github.com/gin-gonic/gin"
	"github.com/mhakimi-adi/munch-bot/pkg/controller"
)

func Item(g *gin.RouterGroup, c *controller.Controller) {
	g.GET("", c.GetItems)
}
