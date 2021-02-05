package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/tidwall/gjson"
)

func RelatedProductsSearch(client string, related []gjson.Result) (gjson.Result, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	shouldQuery := []map[string]interface{}{}

	for _, v := range related {
		match := map[string]interface{}{
			"match": map[string]interface{}{
				"id": v.String(),
			},
		}
		shouldQuery = append(shouldQuery, match)
	}

	//	fmt.Println(shouldQuery)

	boolQuery := map[string]interface{}{
		"should": shouldQuery,
	}
	queryJSON := map[string]interface{}{
		"bool": boolQuery,
	}
	RelatedProducts := map[string]interface{}{
		"query": queryJSON,
	}
	//fmt.Println(RelatedProducts)

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(RelatedProducts); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(client+"_catalog"),
		es.Search.WithBody(&buf),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	jsonBodyRelated := read(res.Body)
	//log.Printf(jsonBodyRelated)
	return gjson.Get(jsonBodyRelated, "hits.hits.#._source"), nil
}

func OneProduct(client string, id string) (OneProductResponse, error) {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.GetSource(client+"_catalog", id, es.GetSource.WithPretty())

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	jsonBody := read(res.Body)

	relatedP := gjson.Get(jsonBody, "related_products")
	related := relatedP.Array()

	hits := len(related)

	var relatedProductGjson gjson.Result
	var OneProductResponse OneProductResponse

	if hits > 0 {
		relatedProductGjson, err = RelatedProductsSearch(client, related)
		if err != nil {
			log.Printf("error:%s", err)
			return OneProductResponse, err
		}
	}

	err = json.Unmarshal([]byte(relatedProductGjson.Raw), &OneProductResponse.RelatedProducts.Products)

	err = json.Unmarshal([]byte(jsonBody), &OneProductResponse.Product)

	err = json.Unmarshal([]byte(strconv.Itoa(hits)), &OneProductResponse.RelatedProducts.Hits)

	if err != nil {
		log.Printf("error:%s", err)
		return OneProductResponse, err
	}

	return OneProductResponse, nil
}
