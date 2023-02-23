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

	conn := postgres.CreateConnection()
	postgres.CreateTable(conn)

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content":  "This is an index page...",
			"itemList": postgres.GetAllItems(conn),
			"newItems": postgres.GetNewItems(conn),
			"oldItems": postgres.GetOldItems(conn),
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
		postgres.InsertItem(name, note, amount, conn)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.GET("item/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.HTML(http.StatusOK, "details.html", gin.H{
			"content": "This is an about page...",
			"item":    postgres.GetItem(id, conn),
		})
	})

	router.POST("/item/:ID/update", func(c *gin.Context) {
		//id := c.PostForm("ID")
		id := c.Params.ByName("ID")
		postgres.UpdateItemStatus(id, conn)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.POST("/item/:ID/delete", func(c *gin.Context) {
		//id := c.PostForm("ID")
		id := c.Params.ByName("ID")
		postgres.DeleteItemStatus(id, conn)
		c.Redirect(http.StatusMovedPermanently, "/manage")
	})

	router.GET("/manage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "manage.html", gin.H{
			"content":      "This is an index page...",
			"itemList":     postgres.GetAllItems(conn),
			"newItems":     postgres.GetNewItems(conn),
			"oldItems":     postgres.GetOldItems(conn),
			"deletedItems": postgres.GetDeletedItems(conn),
		})
	})

	router.Run("localhost:8080")
}
