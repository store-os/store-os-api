package autocomplete

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/tidwall/gjson"
)

type Suggestion struct {
	Title string `json:"title"`
}

type Suggestions []Suggestion

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
