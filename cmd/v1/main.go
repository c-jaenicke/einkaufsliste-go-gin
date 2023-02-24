package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"shopping-list/pkg/postgres"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("asd")

	// := postgres.CreateConnection()
	postgres.CreateConnection()
	postgres.CreateTable()

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	// index page, list of items to buy and old items
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"newItems": postgres.GetItems("new"),
			"oldItems": postgres.GetItems("old"),
		})
	})

	// form for creating a new item
	router.GET("/item/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{
			"title": "Neuen Eintrag anlegen",
		})
	})

	// post path for creating the new item in form
	router.POST("/item/new", func(c *gin.Context) {
		name := c.PostForm("name")
		note := c.PostForm("note")
		amount, _ := strconv.Atoi(c.PostForm("amount"))
		fmt.Println(name, note, amount)
		postgres.InsertItem(name, note, amount)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// show details about an item
	router.GET("item/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.HTML(http.StatusOK, "details.html", gin.H{
			"content": "This is an about page...",
			"item":    postgres.GetItem(id),
		})
	})

	// post path for updating the status of an item
	router.POST("/item/:ID/update", func(c *gin.Context) {
		id := c.Params.ByName("ID")
		postgres.UpdateItemStatus(id)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// path for manage page
	router.GET("/manage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "manage.html", gin.H{
			"itemList":     postgres.GetItems("%"),
			"newItems":     postgres.GetItems("new"),
			"oldItems":     postgres.GetItems("old"),
			"deletedItems": postgres.GetItems("deleted"),
		})
	})

	// post path for changing an item status to deleted
	router.POST("/item/:ID/delete", func(c *gin.Context) {
		id := c.Params.ByName("ID")
		postgres.DeleteItemStatus(id)
		c.Redirect(http.StatusMovedPermanently, "/manage")
	})

	router.Run("localhost:8080")
}
