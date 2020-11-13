package api

import (
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func OneProduct(id string) (Product, error) {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.GetSource("index", id, es.GetSource.WithPretty())

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	jsonBody := read(res.Body)
	//log.Printf(jsonBody)

	var product Product
	err = json.Unmarshal([]byte(jsonBody), &product)
	//log.Printf("%+v", product)
	if err != nil {
		log.Printf("error:%s", err)
		return product, err
	}

	return product, nil
}
