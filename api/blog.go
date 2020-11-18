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

func ListPosts(page string) (Posts, int64, error) {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Build the request body.
	var buf bytes.Buffer

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Fatalf("Error converting page", err)
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
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
		es.Search.WithIndex("index_blog"),
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
	var posts Posts
	err = json.Unmarshal([]byte(raw), &posts)

	hitsResult := gjson.Get(jsonBody, "hits.total.value")

	hits := hitsResult.Int()

	if err != nil {
		log.Printf("error:%s", err)
		return nil, 0, err
	}
	//log.Printf("%+v", posts)

	return posts, hits, nil
}

func OnePost(id string) (Post, error) {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.GetSource("index_blog", id, es.GetSource.WithPretty())

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	jsonBody := read(res.Body)
	//log.Printf(jsonBody)

	var post Post
	err = json.Unmarshal([]byte(jsonBody), &post)
	//log.Printf("%+v", post)
	if err != nil {
		log.Printf("error:%s", err)
		return post, err
	}

	return post, nil
}