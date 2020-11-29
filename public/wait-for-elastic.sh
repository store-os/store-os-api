#!/bin/sh
# wait-for-elastic.sh

set -e
  
host="elasticsearch"

until curl http://$host:9200; do
  >&2 echo "Elastic is unavailable - sleeping"
  sleep 1
done
  
>&2 echo "Elastic is up - executing command"

curl -H "Content-Type: application/json" -XPUT "http://$host:9200/_template/search_template" --data-binary "@/work/public/search_template.json"

curl -H "Content-Type: application/x-ndjson" -XPOST "http://$host:9200/index_search/_bulk" --data-binary "@/work/public/test_search.json"

curl -H "Content-Type: application/x-ndjson" -XPOST "http://$host:9200/index_blog/_bulk" --data-binary "@/work/public/test_blog.json"