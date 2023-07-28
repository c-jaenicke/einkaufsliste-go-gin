package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"shopping-list/ent"
	"shopping-list/pkg/logging"
	"shopping-list/pkg/queries"
)

func main() {
	client, err := ent.Open("postgres", "host=172.22.0.2 port=5432 user=user dbname=db password=pass sslmode=disable")
	if err != nil {
		logging.LogPanic("failed to connect to db: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		logging.LogPanic("failed to create schema: %v", err)
	}

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("items/all", func(c *gin.Context) {
		items, _ := queries.GetAllItems(context.Background(), client)
		c.IndentedJSON(http.StatusOK, items)
	})

	router.POST("item/new", func(c *gin.Context) {
		var itemPost queries.ItemStruct
		if err := c.BindJSON(&itemPost); err != nil {
			logging.LogError("failed to create new item", err)
		}
		err := itemPost.CreateItem(context.Background(), client)
		if err != nil {
			logging.LogError("", err)
		}
	})

	router.GET("stores/all", func(c *gin.Context) {
		stores, _ := queries.GetAllStores(context.Background(), client)
		c.IndentedJSON(http.StatusOK, stores)
	})

	router.POST("store/new", func(c *gin.Context) {
		var storePost queries.StoreStruct
		if err := c.BindJSON(&storePost); err != nil {
			logging.LogError("failed to create new item", err)
		}
		err := storePost.CreateStore(context.Background(), client)
		if err != nil {
			logging.LogError("", err)
		}
	})

	initStores(client)

	router.Run(":8080")
}

// iniStores create a "keiner" if it doesnt already exist
func initStores(client *ent.Client) {
	_, err := queries.GetStoreByName(context.Background(), client, "keiner")
	if err != nil {
		var store = queries.StoreStruct{Name: "keiner"}
		store.CreateStore(context.Background(), client)
	}
}
