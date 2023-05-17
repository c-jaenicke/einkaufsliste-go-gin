package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"shopping-list/pkg/category"
	"shopping-list/pkg/item"
	"shopping-list/pkg/logging"
	"shopping-list/pkg/postgres"
	"strconv"
	"strings"
)

func main() {
	logging.LogInfo("########## Starting app")

	postgres.CreateConnection()
	postgres.CreateTable()

	// uncomment line to switch to release mode
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static("/assets", "./assets")
	router.Static("/images", "./images")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.SetFuncMap(template.FuncMap{
		"GetItemCount":    GetItemCount,
		"GetItemColor":    GetItemColor,
		"GetCategoryName": GetCategoryName,
	})

	// TODO  get all items
	router.GET("/items/new", func(c *gin.Context) {
		items := postgres.GetItems("new")
		c.IndentedJSON(http.StatusOK, items)
	})

	router.POST("item/new", func(c *gin.Context) {
		var item item.Item
		if err := c.BindJSON(&item); err != nil {
			logging.LogError("error creating new item", err)
			c.IndentedJSON(http.StatusBadRequest, nil)
		}

		// TODO improve this, move function to postgres module InsertItemObj
		postgres.InsertItem(item.Name, item.Note, item.Amount, strconv.Itoa(item.Cat_id))
		c.IndentedJSON(http.StatusCreated, item)
	})

	// TODO  post a new item
	// TODO  put an existing item
	// TODO  delete an item
	// TODO get all categories
	// TODO post a new category
	// TODO put an existing category
	// TODO delete a category

	// ALL ROUTES MUST BE ABOVE HERE
	// START GIN
	//
	logging.LogInfo("##### Starting gin on port 8080")
	router.Run(":8080")
}

// GetItemCount counts items that belong to a given category and match the status new.
// Used in categories.html .
// Mapped by router.SetFuncMap()
func GetItemCount(c category.Category) int {
	itemSlice := postgres.GetItemsInCategory(strconv.Itoa(c.ID), "new")
	return len(itemSlice)
}

func GetItemColor(i item.Item) string {
	category := postgres.GetCategory(strconv.Itoa(i.Cat_id))
	return category.Color
}

func GetCategoryName(i item.Item) string {
	id := strconv.Itoa(i.Cat_id)
	return postgres.GetCategory(id).Name

}
