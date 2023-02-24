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

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"itemList":     postgres.GetItems("%"),
			"newItems":     postgres.GetItems("new"),
			"oldItems":     postgres.GetItems("old"),
			"deletedItems": postgres.GetItems("deleted"),
		})
	})

	router.GET("/item/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{
			"content": "This is an about page...",
		})
	})

	router.POST("/item/new", func(c *gin.Context) {
		name := c.PostForm("name")
		note := c.PostForm("note")
		amount, _ := strconv.Atoi(c.PostForm("amount"))
		fmt.Println(name, note, amount)
		postgres.InsertItem(name, note, amount)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.GET("item/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.HTML(http.StatusOK, "details.html", gin.H{
			"content": "This is an about page...",
			"item":    postgres.GetItem(id),
		})
	})

	router.POST("/item/:ID/update", func(c *gin.Context) {
		//id := c.PostForm("ID")
		id := c.Params.ByName("ID")
		postgres.UpdateItemStatus(id)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.POST("/item/:ID/delete", func(c *gin.Context) {
		//id := c.PostForm("ID")
		id := c.Params.ByName("ID")
		postgres.DeleteItemStatus(id)
		c.Redirect(http.StatusMovedPermanently, "/manage")
	})

	router.GET("/manage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "manage.html", gin.H{
			"itemList":     postgres.GetItems("%"),
			"newItems":     postgres.GetItems("new"),
			"oldItems":     postgres.GetItems("old"),
			"deletedItems": postgres.GetItems("deleted"),
		})
	})

	router.Run("localhost:8080")
}
