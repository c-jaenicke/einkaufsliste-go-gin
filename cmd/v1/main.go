package main

import (
	"fmt"
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
	//postgres.CreateTable()

	// uncomment line to switch to release mode
	//gin.SetMode(gin.ReleaseMode)
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
	router.LoadHTMLGlob("./templates/*.html")

	//
	// INDEX PAGE
	// List of new items and old items
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"newItems":     postgres.GetItems("new"),
			"oldItems":     postgres.GetItems("old"),
			"categoryList": postgres.GetAllCategories(),
			"testColor":    "#000000",
		})
	})

	//
	// NEW ITEM FORM
	// Form for creating a new item
	router.GET("/item/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "item-form.html", gin.H{
			"title":        "Neuen Eintrag anlegen",
			"categoryList": postgres.GetAllCategories(),
		})
	})

	router.POST("/item/new", func(c *gin.Context) {
		name := c.PostForm("name")
		note := c.PostForm("note")
		amount, _ := strconv.Atoi(c.PostForm("amount"))
		cat_id := c.PostForm("category")
		postgres.InsertItem(name, note, amount, cat_id)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	//
	// CHANGE ITEM
	// Edit Name, Note and Category of an item
	router.GET("item/:id/change", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.HTML(http.StatusOK, "item-form.html", gin.H{
			"title":        "Artikel Bearbeiten",
			"item":         postgres.GetItem(id),
			"categoryList": postgres.GetAllCategories(),
		})
	})

	router.POST("item/:id/change", func(c *gin.Context) {
		id := c.Params.ByName("id")
		name := c.PostForm("name")
		note := c.PostForm("note")
		amount, _ := strconv.Atoi(c.PostForm("amount"))
		cat_id := c.PostForm("category")
		fmt.Println(cat_id)
		postgres.ChangeItem(id, name, note, amount, cat_id)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	//
	// UPDATE ITEM STATUS
	// Updates the status of an item from new to old and new to old
	router.POST("/item/:id/update", func(c *gin.Context) {
		id := c.Params.ByName("id")
		postgres.UpdateItemStatus(id)
		fmt.Println(c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	//
	// MANAGE
	// Page for bulk deleting items
	router.GET("/manage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "manage.html", gin.H{
			"itemList":     postgres.GetItems("%"),
			"newItems":     postgres.GetItems("new"),
			"oldItems":     postgres.GetItems("old"),
			"deletedItems": postgres.GetItems("deleted"),
			"categoryList": postgres.GetAllCategories(),
		})
	})

	//
	// DELETE ITEM
	// Delete an existing item, only sets status to "deleted"
	router.POST("/item/:id/delete", func(c *gin.Context) {
		id := c.Params.ByName("id")
		postgres.DeleteItemStatus(id)
		c.Redirect(http.StatusMovedPermanently, "/manage")
	})

	//
	// CREATE NEW CATEGORY
	// Form for creating a new category
	router.GET("/category/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "category-form.html", gin.H{
			"title": "Neue Kategorie anlegen",
		})
	})

	router.POST("/category/new", func(c *gin.Context) {
		name := c.PostForm("name")
		color := c.PostForm("color")
		colorName := category.ColorMap[color]
		postgres.CreateCategory(name, color, colorName)
		c.Redirect(http.StatusMovedPermanently, "/category")
	})

	//
	// CATEGORY OVERVIEW
	// Overview of categories
	router.GET("/category", func(c *gin.Context) {
		c.HTML(http.StatusOK, "category-list.html", gin.H{
			"categoryList": postgres.GetAllCategories(),
			"itemList":     postgres.GetItems("new"),
		})
	})

	//
	// CHANGE EXISTING CATEGORY
	// Change the name of an existing category
	router.GET("/category/:id/change", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.HTML(http.StatusOK, "category-form.html", gin.H{
			"title":    "Kategorie Bearbeiten",
			"category": postgres.GetCategory(id),
		})
	})

	router.POST("/category/:id/change", func(c *gin.Context) {
		id := c.Params.ByName("id")
		name := c.PostForm("name")
		postgres.ChangeCategory(id, name)
		c.Redirect(http.StatusMovedPermanently, "/category")
	})

	//
	// DELETE CATEGORY
	// Delete an existing category
	router.POST("/category/:id/delete", func(c *gin.Context) {
		id := c.Params.ByName("id")
		postgres.DeleteCategory(id)
		c.Redirect(http.StatusMovedPermanently, "/category")
	})

	//
	// ITEMS IN CATEGORY
	// List all items in a category
	router.GET("/category/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.HTML(http.StatusOK, "category-details.html", gin.H{
			"title":    "Kategorie Bearbeiten",
			"category": postgres.GetCategory(id),
			"newItems": postgres.GetItemsInCategory(id, "new"),
			"oldItems": postgres.GetItemsInCategory(id, "old"),
		})
	})

	//
	// UPDATE ITEM ON CATEGORY DETAILS VIEW
	// Post for editing an item while in the category details view
	router.POST("/category/:id/item/:iid/update", func(c *gin.Context) {
		id := c.Params.ByName("id")
		iid := c.Params.ByName("iid")
		postgres.UpdateItemStatus(iid)
		fmt.Println(c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/category/"+id)
	})

	//
	// DELETE ALL ITEMS
	// NOT REVERSIBLE! Truncates the items table
	router.POST("manage/deleteall", func(c *gin.Context) {
		postgres.DeleteAllItems()
		c.Redirect(http.StatusMovedPermanently, "/manage")
	})

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
