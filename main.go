package main

import (
	"flag"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/store-os/store-os-api/controller"
	_ "github.com/store-os/store-os-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	indexName  string
	numWorkers int
	flushBytes int
	numItems   int
)

func init() {
	flag.StringVar(&indexName, "index", "test", "Index name")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of indexer workers")
	flag.IntVar(&flushBytes, "flush", 5e+6, "Flush threshold in bytes")
	flag.IntVar(&numItems, "count", 10000, "Number of documents to generate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

// @title Swagger Store OS API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host store-api.alchersan.com
// @BasePath /api/v1
// @schemes https

// @query.collection.format multi

func main() {
	log.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
	r := gin.Default()

	r.Use(cors.Default())

	c := controller.NewController()

	v1 := r.Group("/api/v1/:client")
	{
		autocomplete := v1.Group("/autocomplete")
		{
			autocomplete.GET("", c.Autocomplete)
		}
		products := v1.Group("/products")
		{
			products.GET("", c.Search)
			products.GET(":id", c.OneProduct)
			products.POST(":id/update", c.EditProduct)
		}
		aggs := v1.Group("/aggs")
		{
			aggs.GET("", c.Aggs)
		}
		posts := v1.Group("/blog")
		{
			posts.GET("", c.ListPosts)
			posts.GET(":id", c.OnePost)
		}
	}

	health := r.Group("/health")
	{
		health.GET("", func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
			c.Header("Access-Control-Allow-Headers", "x-access-token, Origin, X-Requested-With, Content-Type, Accept")
			c.JSON(200, gin.H{
				"status": "health",
			})
		})
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
