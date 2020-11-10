package main

import (
	"flag"
	"math/rand"
	"runtime"
	"time"

	"github.com/store-os/store-os-api/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

// @host localhost:8080
// @BasePath /api/v1

// @query.collection.format multi

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		search := v1.Group("/search")
		{
			search.GET("", c.Search)
		}
		suggest := v1.Group("/suggest")
		{
			suggest.GET("", c.Autocomplete)
		}
		products := v1.Group("/products")
		{
			products.GET("", c.ListProducts)
			products.GET(":id", c.OneProduct)
		}
		health := v1.Group("/health")
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
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
