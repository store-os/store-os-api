#!/bin/bash
curl -H "Content-Type: application/json" -XPUT "http://172.18.0.1:9200/_template/search_template" --data-binary "@/work/public/search_template.json"

curl -H "Content-Type: application/x-ndjson" -XPOST "http://172.18.0.1:9200/index_search/_bulk" --data-binary "@/work/public/test_search.json"

curl -H "Content-Type: application/x-ndjson" -XPOST "http://172.18.0.1:9200/index_blog/_bulk" --data-binary "@/work/public/test_blog.json"