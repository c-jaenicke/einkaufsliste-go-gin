package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-list/pkg/category"
	"shopping-list/pkg/item"
	"shopping-list/pkg/logging"
	"shopping-list/pkg/postgres"
	"strconv"
)

func main() {
	logging.LogInfo("########## Starting app")

	postgres.CreateConnection()
	postgres.CreateTable()

	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// use DEFAULT config for gin-contrib/cors middleware
	// ATTENTION: THIS ALLOWS ALL ORIGINS
	// see https://github.com/gin-contrib/cors for more information
	router.Use(cors.Default())

	// Get all items with give status
	router.GET("/items/:status", func(c *gin.Context) {
		status := c.Params.ByName("status")
		c.IndentedJSON(http.StatusOK, postgres.GetItemsWithStatus(status))
	})

	// Get all items, regardless of status
	router.GET("/items/all", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, postgres.GetAllItems())
	})

	// Post new item, 201 on successful post, 400 on error
	router.POST("/item/new", func(c *gin.Context) {
		var item item.Item
		if err := c.BindJSON(&item); err != nil {
			logging.LogError("Unable to create new item from body", err)
			c.IndentedJSON(http.StatusBadRequest, nil)
		}

		postgres.SaveItem(item)
		c.IndentedJSON(http.StatusCreated, postgres.GetAllItems())
	})

	// Put an existing item
	router.PUT("/item/:id/update", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			logging.LogError("Failed to convert ID to int", err)
			c.IndentedJSON(http.StatusBadRequest, nil)
		}

		var item item.Item
		if err := c.BindJSON(&item); err != nil {
			logging.LogError("Error updating an item", err)
			c.IndentedJSON(http.StatusBadRequest, nil)
		}

		postgres.UpdateItem(item, id)
		c.IndentedJSON(http.StatusOK, item)
	})

	// Change item status from old to new
	router.POST("/item/:id/switch", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			logging.LogError("Failed to convert ID to int", err)
			c.IndentedJSON(http.StatusBadRequest, nil)
		}
		postgres.SwitchItemStatus(id)
		c.IndentedJSON(http.StatusOK, postgres.GetAllItems())
	})

	// Delete an item
	router.DELETE("/item/:id/delete", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			logging.LogError("Failed to convert ID to int", err)
			c.IndentedJSON(http.StatusBadRequest, nil)
		}
		postgres.DeleteItem(id)
		c.IndentedJSON(http.StatusOK, postgres.GetAllItems())
	})

	// Get all categories
	router.GET("/categories", func(c *gin.Context) {
		categories := postgres.GetAllCategories()
		c.IndentedJSON(http.StatusOK, categories)
	})

	// TODO get all categories
	// Get slice of all categories

	// TODO post a new category
	// Post create a new category

	// TODO put an existing category
	// Update an existing category

	// TODO delete a category
	// Delete a category

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
