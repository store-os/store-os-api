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

func getAggs() map[string]interface{} {
	availableAggs := aggsVariable("available", 2)
	categoriesAggs := aggsVariable("category.keyword", 30)
	subcategoriesAggs := aggsVariable("subcategory.keyword", 30)
	brandAggs := aggsVariable("brand.keyword", 10)
	genderAggs := aggsVariable("gender.keyword", 3)

	priceRangeAggs := map[string]interface{}{
		"range": map[string]interface{}{
			"field": "price",
			"ranges": []map[string]interface{}{
				{"from": 0.0, "to": 100.0, "key": "0 - 100"},
				{"from": 101.0, "to": 500.0, "key": "101 - 500"},
				{"from": 501.0, "to": 2500.0, "key": "501 - 2500"},
				{"from": 2501.0, "to": 5000.0, "key": "2501 - 5000"},
				{"from": 5001.0, "to": 10000.0, "key": "5001 - 10000"},
				{"from": 100001.0, "to": 10000000.0, "key": "100001+"},
			},
		},
	}

	aggs := map[string]interface{}{
		"available":     availableAggs,
		"categories":    categoriesAggs,
		"subcategories": subcategoriesAggs,
		"brand":         brandAggs,
		"gender":        genderAggs,
		"price":         priceRangeAggs,
	}
	return aggs
}

func aggsVariable(field string, size int) map[string]interface{} {
	aggsQuery := map[string]interface{}{
		"terms": map[string]interface{}{
			"field": field,
			"size":  size,
		},
	}
	return aggsQuery
}

func filtersVariable(field string, value string) map[string]interface{} {
	query := map[string]interface{}{
		"term": map[string]interface{}{
			field: map[string]interface{}{
				"value": value,
			},
		},
	}
	return query
}

func matchAll() map[string]interface{} {
	query := map[string]interface{}{
		"match_all": map[string]interface{}{},
	}
	return query
}

func filtering(category string, subcategory string, q string) []map[string]interface{} {
	query := []map[string]interface{}{}

	subMap := map[string]interface{}{}
	if category != "" {
		subMap = filtersVariable("category.keyword", category)
		query = append(query, subMap)
	}

	subMap = nil
	if subcategory != "" {
		subMap = filtersVariable("subcategory.keyword", subcategory)
		query = append(query, subMap)
	}
	subMap = nil
	if q == "" {
		subMap = matchAll()
		query = append(query, subMap)
	}

	return query
}

func getQuery(q string, category string, subcategory string) map[string]interface{} {

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

	should := map[string]interface{}{
		"should": []map[string]interface{}{
			desc,
			ref,
			tit,
		},
	}

	filters := map[string]interface{}{
		"must": filtering(category, subcategory, q),
	}

	for k, v := range should {
		filters[k] = v
	}

	queryJSON := map[string]interface{}{}

	boolQuery := map[string]interface{}{}

	if (category != "") || (subcategory != "") || (q == "") {
		boolQuery = filters
	} else {
		boolQuery = should
	}
	queryJSON = map[string]interface{}{
		"bool": boolQuery,
	}

	return queryJSON
}

func Search(q string, page string, category string, subcategory string) (SearchResponse, error) {

	es, err := elasticsearch.NewDefaultClient()

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Build the request body.
	var buf bytes.Buffer

	aggs := getAggs()
	queryJSON := getQuery(q, category, subcategory)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Fatalf("Error converting page", err)
	}

	query := map[string]interface{}{
		"query": queryJSON,
		"aggs":  aggs,
		"from":  pageInt * size,
		"size":  size,
	}

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

	aggregations := gjson.Get(jsonBody, "aggregations")

	hitsResult := gjson.Get(jsonBody, "hits.total.value")

	var jsonByte []byte
	var raw []byte
	if result.Index > 0 {
		raw = jsonByte[result.Index : result.Index+len(result.Raw)]
	} else {
		raw = []byte(result.Raw)
	}
	var searchResponse SearchResponse

	err = json.Unmarshal([]byte(aggregations.Raw), &searchResponse.Aggregations)

	err = json.Unmarshal([]byte(raw), &searchResponse.Products)

	err = json.Unmarshal([]byte(hitsResult.Raw), &searchResponse.Hits)

	if err != nil {
		log.Printf("error:%s", err)
		return searchResponse, err
	}
	//log.Printf("%+v", searchResponse)

	return searchResponse, nil
}
