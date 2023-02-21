package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhakimi-adi/munch-bot/pkg/model"
	util "github.com/mhakimi-adi/munch-bot/pkg/util"
)

// ListItems godoc
// @Summary      List all items
// @Description  get items
// @Tags         item
// @Accept       json
// @Produce      json
// @Success      200  {array}  model.Item
// @Failure	     400  {object} util.HTTPError
// @Failure	     400  {object} util.HTTPError
// @Failure	     500  {object} util.HTTPError
// @Router       /item  [get]
func (c *Controller) GetItems(ctx *gin.Context) {
	items, err := model.GetItems()
	if err != nil {
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.IndentedJSON(http.StatusAccepted, items)
}
