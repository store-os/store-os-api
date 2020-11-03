package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func read(r io.Reader) string {
	var b bytes.Buffer
	b.ReadFrom(r)
	return b.String()
}

type Product struct {
	Title         string `json:"title"`
	Category      string `json:"category"`
	Image         string `json:"image"`
	Price         int    `json:"price"`
	DiscountPrice int    `json:"discount_price"`
	Link          string `json:"link"`
}

type Products []Product

type Suggestion struct {
	Title string `json:"title"`
}

type Suggestions []Suggestion

var (
	indexName  string
	numWorkers int
	flushBytes int
	numItems   int
)

func Search(q string) Products {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Build the request body.
	var buf bytes.Buffer

	log.Printf(q)

	desc := map[string]interface{}{
		"match_phrase": map[string]interface{}{
			"description": map[string]interface{}{
				"query": q,
			},
		},
	}

	ref := map[string]interface{}{
		"match": map[string]interface{}{
			"reference": q,
		},
	}

	tit := map[string]interface{}{
		"multi_match": map[string]interface{}{
			"query": q,
			"type":  "bool_prefix",
			"fields": []string{
				"title",
				"title._2gram",
				"title._3gram",
			},
		},
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					desc,
					ref,
					tit,
				},
			},
		},
		"size": "100",
	}

	//fmt.Println(query)

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("index"),
		es.Search.WithBody(&buf),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	jsonBody := read(res.Body)
	//log.Printf(jsonBody)

	result := gjson.Get(jsonBody, "hits.hits.#._source")

	var jsonByte []byte
	var raw []byte
	if result.Index > 0 {
		raw = jsonByte[result.Index : result.Index+len(result.Raw)]
	} else {
		raw = []byte(result.Raw)
	}
	var products Products
	err = json.Unmarshal([]byte(raw), &products)

	if err != nil {
		log.Printf("error:%s", err)
	}
	//log.Printf("%+v", products)

	return products

}

func Autocomplete(q string) Suggestions {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Build the request body.
	var buf bytes.Buffer

	log.Printf(q)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query": q,
				"type":  "bool_prefix",
				"fields": []string{
					"title",
					"title._2gram",
					"title._3gram",
				},
			},
		},
		"size": "3",
	}

	//fmt.Println(query)

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("index"),
		es.Search.WithBody(&buf),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	jsonBody := read(res.Body)
	//log.Printf(jsonBody)

	result := gjson.Get(jsonBody, "hits.hits.#._source")

	var jsonByte []byte
	var raw []byte
	if result.Index > 0 {
		raw = jsonByte[result.Index : result.Index+len(result.Raw)]
	} else {
		raw = []byte(result.Raw)
	}
	var suggestions Suggestions
	err = json.Unmarshal([]byte(raw), &suggestions)

	if err != nil {
		log.Printf("error:%s", err)
	}
	//log.Printf("%+v", products)

	return suggestions

}

func init() {
	flag.StringVar(&indexName, "index", "test", "Index name")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of indexer workers")
	flag.IntVar(&flushBytes, "flush", 5e+6, "Flush threshold in bytes")
	flag.IntVar(&numItems, "count", 10000, "Number of documents to generate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	// This handler will match /search/:query but will not match /search/ or /search
	r.GET("/search", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "x-access-token, Origin, X-Requested-With, Content-Type, Accept")
		query := c.Query("q")
		body := Search(query)
		c.JSON(200, gin.H{
			"products": body,
		})
	})

	r.GET("/suggest", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "x-access-token, Origin, X-Requested-With, Content-Type, Accept")
		query := c.Query("q")
		body := Autocomplete(query)
		c.JSON(200, gin.H{
			"suggestions": body,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "x-access-token, Origin, X-Requested-With, Content-Type, Accept")
		c.JSON(200, gin.H{
			"status": "health",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
