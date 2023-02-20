package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"

	"strconv"
)

type Item struct {
	ItemId      int    `json:"item_id"`
	ItemName    string `json:"item_name"`
	ItemDesc    string `json:"item_desc"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	IsDeleted   int    `json:"is_deleted"`
}

var items = []Item{
	{
		ItemId:      1,
		ItemName:    "Oil",
		ItemDesc:    "Oil",
		CreatedDate: "2020-01-01",
		UpdatedDate: "2020-01-01",
		IsDeleted:   0,
	},
	{
		ItemId:      2,
		ItemName:    "Banana",
		ItemDesc:    "Banana",
		CreatedDate: "2020-01-01",
		UpdatedDate: "2020-01-01",
		IsDeleted:   0,
	},
	{
		ItemId:      3,
		ItemName:    "Orange",
		ItemDesc:    "Orange",
		CreatedDate: "2020-01-01",
		UpdatedDate: "2020-01-01",
		IsDeleted:   0,
	},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, items)
}

func createItem(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func itemById(c *gin.Context) {
	id := c.Param("item_id")
	idInt, _ := strconv.Atoi(id)
	item, err := getItemById(idInt)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item NOT Found"})
		return
	}

	c.IndentedJSON(http.StatusFound, item)
}

func getItemById(id int) (*Item, error) {
	for idx, item := range items {
		if item.ItemId == id {
			return &items[idx], nil
		}
	}

	return nil, errors.New("Item NOT found")
}

func main() {
	router := gin.Default()
	router.GET("/item", getItems)
	router.GET("/item/:item_id", itemById)
	router.POST("/item", createItem)
	router.Run("localhost:8080")
}
