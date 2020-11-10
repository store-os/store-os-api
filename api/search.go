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

func Search(q string, page string) (Products, int64, error) {

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
			"id": q,
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

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Fatalf("Error converting page", err)
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
		"from": pageInt * size,
		"size": size,
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

	hitsResult := gjson.Get(jsonBody, "hits.total.value")

	hits := hitsResult.Int()

	if err != nil {
		log.Printf("error:%s", err)
		return nil, 0, err
	}
	//log.Printf("%+v", products)

	return products, hits, nil
}
