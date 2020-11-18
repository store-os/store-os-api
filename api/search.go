package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/tidwall/gjson"
)

func getAggs() map[string]interface{} {
	availableAggs := aggsVariable("available", 2)
	categoriesAggs := composeAggsVariable("levels.category.keyword", "levels.subcategory.keyword", "levels.subsubcategory.keyword", 30)

	brandAggs := aggsVariable("brand.keyword", 10)
	genderAggs := aggsVariable("gender.keyword", 3)
	minPrice := map[string]interface{}{
		"min": map[string]interface{}{
			"field": "price",
		},
	}
	maxPrice := map[string]interface{}{
		"min": map[string]interface{}{
			"field": "price",
		},
	}

	aggs := map[string]interface{}{
		"available":  availableAggs,
		"categories": categoriesAggs,
		"brand":      brandAggs,
		"gender":     genderAggs,
		"minPrice":   minPrice,
		"maxPrice":   maxPrice,
	}
	return aggs
}

func composeAggsVariable(category string, subcategory string, subsubcategory string, size int) map[string]interface{} {
	composeAggsVariable := map[string]interface{}{
		"terms": map[string]interface{}{
			"field": category,
			"size":  size,
		},
		"aggs": map[string]interface{}{
			"subcategories": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": subcategory,
					"size":  size,
				},
				"aggs": map[string]interface{}{
					"subsubcategories": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": subsubcategory,
							"size":  size,
						},
					},
				},
			},
		},
	}
	return composeAggsVariable
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

func filtering(category string, subcategory string, subsubcategory string, from string, to string, q string) []map[string]interface{} {
	query := []map[string]interface{}{}

	subMap := map[string]interface{}{}
	if category != "" {
		subMap = filtersVariable("levels.category.keyword", category)
		query = append(query, subMap)
	}

	subMap = nil
	if subcategory != "" {
		subMap = filtersVariable("levels.subcategory.keyword", subcategory)
		query = append(query, subMap)
	}
	subMap = nil
	if subsubcategory != "" {
		subMap = filtersVariable("levels.subsubcategory.keyword", subcategory)
		query = append(query, subMap)
	}
	if (from != "") && (to != "") {
		rangeQuery := map[string]interface{}{
			"range": map[string]interface{}{
				"price": map[string]interface{}{
					"gte": from,
					"lte": to,
				},
			},
		}

		query = append(query, rangeQuery)
	}
	subMap = nil
	if q == "" {
		subMap = matchAll()
		query = append(query, subMap)
	}

	return query
}

func getSort(field string, order string) []map[string]interface{} {

	sortShort := map[string]interface{}{
		field: map[string]interface{}{
			"order": order,
		},
	}
	sort := []map[string]interface{}{
		sortShort,
	}
	return sort
}

func getQuery(q string, category string, subcategory string, subsubcategory string, from string, to string) map[string]interface{} {

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
		"must": filtering(category, subcategory, subsubcategory, from, to, q),
	}

	for k, v := range should {
		filters[k] = v
	}

	queryJSON := map[string]interface{}{}

	boolQuery := map[string]interface{}{}

	if (category != "") || (subcategory != "") || (q == "") || (subsubcategory != "") || (from != "") || (to != "") {
		boolQuery = filters
	} else {
		boolQuery = should
	}

	queryJSON = map[string]interface{}{
		"bool": boolQuery,
	}

	return queryJSON
}

func Search(q string, page string, category string, subcategory string, subsubcategory string, fieldSort string, order string, from string, to string) (SearchResponse, error) {

	es, err := elasticsearch.NewDefaultClient()

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Build the request body.
	var buf bytes.Buffer

	aggs := getAggs()
	queryJSON := getQuery(q, category, subcategory, subsubcategory, from, to)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Fatalf("Error converting page", err)
	}

	sort := getSort(fieldSort, order)

	query := map[string]interface{}{
		"query": queryJSON,
		"aggs":  aggs,
		"from":  pageInt * size,
		"size":  size,
		"sort":  sort,
	}
	fmt.Println(query)
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
