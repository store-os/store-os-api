package api

import (
	"encoding/json"
	"log"
	"strings"
	"github.com/elastic/go-elasticsearch/v8"
)

func EditProduct(client string, id string, product Product) (Product, error) {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	var updateProduct UpdateProduct
	updateProduct.Doc = product
	e, err := json.Marshal(updateProduct)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	
	res, err := es.Update(client+"_catalog", id, strings.NewReader(string(e)))

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	jsonBody := read(res.Body)
	log.Printf(jsonBody)

	err = json.Unmarshal([]byte(jsonBody), &product)
	//log.Printf("%+v", product)
	if err != nil {
		log.Printf("error:%s", err)
		return product, err
	}

	return product, nil
}
