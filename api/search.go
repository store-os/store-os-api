package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

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
			"field": "final_price",
		},
	}
	maxPrice := map[string]interface{}{
		"max": map[string]interface{}{
			"field": "final_price",
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

func filtersVariable(field string, value []string) map[string]interface{} {
	query := map[string]interface{}{
		"terms": map[string]interface{}{
			field: value,
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

func filtering(category []string, subcategory []string, subsubcategory []string, from string, to string, q string) []map[string]interface{} {
	query := []map[string]interface{}{}

	subMap := map[string]interface{}{}

	if category != nil {
		subMap = filtersVariable("levels.category.keyword", category)
		query = append(query, subMap)
	}

	subMap = nil
	if subcategory != nil {
		subMap = filtersVariable("levels.subcategory.keyword", subcategory)
		query = append(query, subMap)
	}

	subMap = nil
	if subsubcategory != nil {
		subMap = filtersVariable("levels.subsubcategory.keyword", subsubcategory)
		query = append(query, subMap)
	}

	if (from != "") && (to != "") {
		rangeQuery := map[string]interface{}{
			"range": map[string]interface{}{
				"final_price": map[string]interface{}{
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

func getQuery(q string, category []string, subcategory []string, subsubcategory []string, from string, to string) map[string]interface{} {

	miniDescription := map[string]interface{}{
		"simple_query_string": map[string]interface{}{
			"query":              "\"" + q + "\"",
			"quote_field_suffix": ".exact",
			"fields": []string{
				"mini_description",
				"title",
			},
		},
	}

	ref := map[string]interface{}{
		"match": map[string]interface{}{
			"id": q,
		},
	}

	countSpace := strings.Count(q, " ")
	//fmt.Println("Number of spaces:", countSpace)
	should := map[string]interface{}{}
	if countSpace == 0 {
		tit := map[string]interface{}{
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
		should = map[string]interface{}{
			"should": []map[string]interface{}{
				miniDescription,
				ref,
				tit,
			},
		}
	} else {
		should = map[string]interface{}{
			"should": []map[string]interface{}{
				miniDescription,
				ref,
			},
		}
	}

	filters := map[string]interface{}{
		"must": filtering(category, subcategory, subsubcategory, from, to, q),
	}

	for k, v := range should {
		filters[k] = v
	}

	queryJSON := map[string]interface{}{}

	boolQuery := map[string]interface{}{}

	if (category != nil) || (subcategory != nil) || (q == "") || (subsubcategory != nil) || (from != "") || (to != "") {
		boolQuery = filters
	} else {
		boolQuery = should
	}

	queryJSON = map[string]interface{}{
		"bool": boolQuery,
	}

	return queryJSON
}

func Search(client string, q string, page string, category []string, subcategory []string, subsubcategory []string, fieldSort string, order string, from string, to string, size string) (SearchResponse, error) {

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

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		log.Fatalf("Error converting size", err)
	}

	sort := getSort(fieldSort, order)

	query := map[string]interface{}{
		"query": queryJSON,
		"aggs":  aggs,
		"from":  pageInt * sizeInt,
		"size":  sizeInt,
		"sort":  sort,
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
