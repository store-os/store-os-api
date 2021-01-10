package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/tidwall/gjson"
)

func SearchAutocomplete(client string, q string, category []string, subcategory []string, subsubcategory []string, from string, to string) (Autocompletes, error) {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	termFilter := filtering(category, subcategory, subsubcategory, from, to, q)
	//fmt.Println(term_filter)
	// Build the request body.
	var buf bytes.Buffer

	log.Printf(q)

	titleAutocomplete := map[string]interface{}{
		"multi_match": map[string]interface{}{
			"query": q,
			"type":  "bool_prefix",
			"fields": []string{
				"title.autocomplete",
				"title.autocomplete._2gram",
				"title.autocomplete._3gram",
			},
		},
	}

	termFilter = append(termFilter, titleAutocomplete)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": termFilter,
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
		es.Search.WithIndex(client+"_catalog"),
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
		return nil, err
	}
	//log.Printf("%+v", products)

	var autocompleteResponse Autocompletes

	for _, value := range products {
		var autocomplete Autocomplete

		autocomplete.Title = value.Title
		autocomplete.Image = value.Images[0]
		autocomplete.ID = value.ID
		autocompleteResponse = append(autocompleteResponse, autocomplete)

	}

	return autocompleteResponse, nil

}
