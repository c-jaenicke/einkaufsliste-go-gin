package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"shopping-list/pkg/logging"
	"shopping-list/pkg/postgres"
	"strconv"
	"strings"
)

func main() {
	logging.LogInfo("########## Starting app")

	// := postgres.CreateConnection()
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
	router.LoadHTMLGlob("./templates/*.html")

	// index page, list of items to buy and old items
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"newItems":  postgres.GetItems("new"),
			"oldItems":  postgres.GetItems("old"),
			"testColor": "#000000",
		})
	})

	// form for creating a new item
	router.GET("/item/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{
			"title":        "Neuen Eintrag anlegen",
			"categoryList": postgres.GetAllCategories(),
		})
	})

	// post path for creating the new item in form
	router.POST("/item/new", func(c *gin.Context) {
		name := c.PostForm("name")
		note := c.PostForm("note")
		amount, _ := strconv.Atoi(c.PostForm("amount"))
		cat_id := c.PostForm("category")
		postgres.InsertItem(name, note, amount, cat_id)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// update an item
	router.GET("item/:id/change", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.HTML(http.StatusOK, "form.html", gin.H{
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
		postgres.ChangeItem(id, name, note, amount)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// post path for updating the status of an item
	router.POST("/item/:id/update", func(c *gin.Context) {
		id := c.Params.ByName("id")
		postgres.UpdateItemStatus(id)
		fmt.Println(c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// path for manage page
	router.GET("/manage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "manage.html", gin.H{
			"itemList":     postgres.GetItems("%"),
			"newItems":     postgres.GetItems("new"),
			"oldItems":     postgres.GetItems("old"),
			"deletedItems": postgres.GetItems("deleted"),
			"categoryList": postgres.GetAllCategories(),
		})
	})

	// post path for changing an item status to deleted
	router.POST("/item/:id/delete", func(c *gin.Context) {
		id := c.Params.ByName("id")
		postgres.DeleteItemStatus(id)
		c.Redirect(http.StatusMovedPermanently, "/manage")
	})

	//
	// CREATE NEW CATEGORY
	//
	router.GET("/category/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "category-form.html", gin.H{
			"title": "Neue Kategorie anlegen",
		})
	})

	router.POST("/category/new", func(c *gin.Context) {
		name := c.PostForm("name")
		postgres.CreateCategory(name)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	//
	// CATEGORY OVERVIEW
	//
	router.GET("/category", func(c *gin.Context) {
		c.HTML(http.StatusOK, "categories.html", gin.H{
			"categoryList": postgres.GetAllCategories(),
		})
	})

	//
	// CHANGE EXISTING CATEGORY
	//
	router.GET("category/:id/change", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.HTML(http.StatusOK, "category-form.html", gin.H{
			"title":    "Kategorie Bearbeiten",
			"category": postgres.GetCategory(id),
		})
	})

	router.POST("category/:id/change", func(c *gin.Context) {
		id := c.Params.ByName("id")
		name := c.PostForm("name")
		postgres.ChangeCategory(id, name)
		c.Redirect(http.StatusMovedPermanently, "/category")
	})

	//
	// ITEMS IN CATEGORY
	//
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
	//
	router.POST("/category/:id/item/:iid/update", func(c *gin.Context) {
		id := c.Params.ByName("id")
		iid := c.Params.ByName("iid")
		postgres.UpdateItemStatus(iid)
		fmt.Println(c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/category/"+id)
	})

	logging.LogInfo("##### Starting gin on port 8080")
	router.Run(":8080")
}
