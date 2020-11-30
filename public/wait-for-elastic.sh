#!/bin/sh
# wait-for-elastic.sh

set -e
  
host="elasticsearch"

until curl http://$host:9200; do
  >&2 echo "Elastic is unavailable - sleeping"
  sleep 1
done
  
>&2 echo "Elastic is up - executing command"

curl -H "Content-Type: application/json" -XPUT "http://$host:9200/_template/catalog_template" --data-binary "@/work/public/catalog_template.json"

curl -H "Content-Type: application/x-ndjson" -XPOST "http://$host:9200/alchersan_catalog/_bulk" --data-binary "@/work/public/alchersan_catalog.json"

curl -H "Content-Type: application/x-ndjson" -XPOST "http://$host:9200/alchersan_blog/_bulk" --data-binary "@/work/public/alchersan_blog.json"