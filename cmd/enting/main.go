package main

import (
	"context"
	"fmt"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent"
	"github.com/c-jaenicke/einkaufsliste-go-gin/pkg/logging"
	"github.com/c-jaenicke/einkaufsliste-go-gin/pkg/queries"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Server struct {
	http    *gin.Engine
	db      *ent.Client
	envKeys *EnvKeys
}

var server Server

func runHttp() {
	// set up router here
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: server.envKeys.AllowedOrigins,
	}))
	server.http = router

	//
	// ITEM MAPPINGS
	//
	router.GET("item/all", func(c *gin.Context) {
		items, _ := queries.GetAllItems(context.Background(), server.db)
		c.IndentedJSON(http.StatusOK, items)
	})
	// TODO: merge both mappings together
	// TODO: improve this mapping, right now you can only filter by one of the queries. This also requires creating a query method in the queries/item.go package to query by multiple attributes
	router.GET("item/specific", func(c *gin.Context) {
		status := c.Query("status")
		what := c.Query("type")
		id := c.Query("id")
		name := c.Query("name")

		var err error
		var items []*ent.Item

		// if status query given -> query by status
		// else if type == store
		// 		id not empty -> query store by id
		//		name not empty -> query store by name
		// else if type == category
		// 		id not empty -> query category by id
		//		name not empty -> query category by name
		// else -> get all items

		if status != "" {
			items, err = queries.GetAllItemsByStatus(context.Background(), server.db, status)
		} else if what == "store" {

			if id != "" {
				storeIdInt, err := strconv.Atoi(id)
				if err != nil {
					logging.LogError("", err)
					c.Status(http.StatusInternalServerError)
				}

				items, err = queries.GetAllItemsByStoreId(context.Background(), server.db, storeIdInt)
			} else if name != "" {
				items, err = queries.GetAllItemsByStoreName(context.Background(), server.db, name)
			}

		} else if what == "category" {

			if id != "" {
				categoryId, err := strconv.Atoi(id)
				if err != nil {
					logging.LogError("", err)
					c.Status(http.StatusInternalServerError)
				}

				items, err = queries.GetAllItemsByCategoryId(context.Background(), server.db, categoryId)
			} else if name != "" {
				items, err = queries.GetAllItemsByCategoryName(context.Background(), server.db, name)
			}

		} else {
			items, _ = queries.GetAllItems(context.Background(), server.db)
		}

		if err != nil {
			c.Status(http.StatusInternalServerError)
		}
		c.IndentedJSON(http.StatusOK, items)
	})

	router.POST("item/new", func(c *gin.Context) {
		var itemPost queries.ItemStruct
		if err := c.BindJSON(&itemPost); err != nil {
			logging.LogError("failed to create new item", err)
		}
		err := itemPost.Create(context.Background(), server.db)
		if err != nil {
			logging.LogError("", err)
		}
	})

	router.GET("item/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}

		item, err := queries.GetItemById(context.Background(), server.db, id)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		c.IndentedJSON(http.StatusOK, item)

	})

	router.DELETE("item/:id/delete", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}

		it := queries.ItemStruct{Id: id}
		err = it.Delete(context.Background(), server.db)
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}
		c.Status(http.StatusOK)
	})

	router.PUT("item/:id/update", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}

		var itemPost queries.ItemStruct
		if err := c.BindJSON(&itemPost); err != nil {
			logging.LogError("failed to create new item", err)
		}

		itemPost.Id = id
		err = itemPost.Update(context.Background(), server.db)
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}
		c.Status(http.StatusOK)
	})

	router.PATCH("item/:id/switch", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}

		err = queries.SwitchItemStatus(context.Background(), server.db, id)
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}
		c.Status(http.StatusOK)
	})

	//
	// STORE MAPPINGS
	//
	router.GET("stores/all", func(c *gin.Context) {
		stores, _ := queries.GetAllStores(context.Background(), server.db)
		c.IndentedJSON(http.StatusOK, stores)
	})

	router.POST("store/new", func(c *gin.Context) {
		var storePost queries.StoreStruct
		if err := c.BindJSON(&storePost); err != nil {
			logging.LogError("failed to create new item", err)
		}
		err := storePost.Create(context.Background(), server.db)
		if err != nil {
			logging.LogError("", err)
		}
	})

	router.DELETE("store/:id/delete", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}

		st := queries.StoreStruct{Id: id}
		err = st.Delete(context.Background(), server.db)
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}
		c.Status(http.StatusOK)
	})

	// TODO: possible update mapping here

	//
	// CATEGORY MAPPING
	//
	router.GET("category/all", func(c *gin.Context) {
		cats, err := queries.GetAllCategories(context.Background(), server.db)
		if err != nil {
			c.Status(http.StatusInternalServerError)
		}
		c.IndentedJSON(http.StatusOK, cats)
	})

	router.POST("category/new", func(c *gin.Context) {
		var cat queries.CategoryStruct
		if err := c.BindJSON(&cat); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		err := cat.Create(context.Background(), server.db)
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}
	})

	router.DELETE("category/:id/delete", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}

		cat := queries.CategoryStruct{Id: id}
		err = cat.Delete(context.Background(), server.db)
		if err != nil {
			logging.LogError("", err)
			c.Status(http.StatusInternalServerError)
		}
		c.Status(http.StatusOK)
	})

	// TODO: possible update mapping here

	err := server.http.Run(":8080")
	if err != nil {
		logging.LogPanic("failed to start http server: %w", err)
	}
}

// initDatabase initialize database connection, create schema and entries
func initDatabase() {
	//dataSourceString := "host=einkaufsliste-db port=5432 user=user dbname=db password=pass sslmode=disable"

	client, err := ent.Open("postgres", server.envKeys.Dss)
	logging.LogInfo("attempting connection to db under: " + server.envKeys.Dss)
	if err != nil {
		logging.LogPanic("failed to connect to db: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		logging.LogPanic("failed to create schema: %v", err)
	}

	server.db = client

	initStores()
	initCategories()
}

// initStores create a "keiner" if it doesn't already exist
func initStores() {
	_, err := queries.GetStoreByName(context.Background(), server.db, "keiner")
	if err != nil {
		var store = queries.StoreStruct{Name: "keiner"}
		errCreate := store.Create(context.Background(), server.db)
		if errCreate != nil {
			logging.LogPanic("failed to create initial category: ", errCreate)
		}
	}
}

// initCategories create a "keine" category if it doesn't exist already
func initCategories() {
	_, err := queries.GetCategoryByName(context.Background(), server.db, "keine")
	if err != nil {
		var cat = queries.CategoryStruct{
			Name:  "keine",
			Color: "#ffffff",
		}
		errCreate := cat.Create(context.Background(), server.db)
		if errCreate != nil {
			logging.LogPanic("failed to create initial category: ", errCreate)
		}
	}
}

type EnvKeys struct {
	Dss            string
	AllowedOrigins []string
}

func getEnvKeys() (*EnvKeys, error) {
	// TODO: rework all of this, i dont even know anymore
	envKeys := &EnvKeys{}
	envFile, err := godotenv.Read(".env")
	if err == nil {
		envKeys.Dss = envFile["DSS"]
		envKeys.AllowedOrigins = strings.Split(envFile["ALLOWED_ORIGINS"], ",")
	}

	if envKeys.Dss == "" {
		envKeys.Dss = os.Getenv("DSS")
	}

	if len(envKeys.AllowedOrigins) == 0 {
		envKeys.AllowedOrigins = strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	}

	logging.LogInfo(fmt.Sprintf("env keys set to: %v", envKeys))
	return envKeys, err
}

func main() {
	var err error
	server.envKeys, err = getEnvKeys()
	if err != nil {
		logging.LogError("error getting env keys:", err)
	}
	initDatabase()
	runHttp()
}
