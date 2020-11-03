docker network create localhost

docker run -d --network localhost --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.1

docker run -d --network localhost -p 9000:9000 --name cerebro yannart/cerebro:latest

docker run -d --link <elastic-search-container-id>:elasticsearch -p 5601:5601 --network localhost docker.elastic.co/kibana/kibana:7.6.1

docker run -d --network localhost -v "$PWD:/work" --name curl curlimages/curl

docker exec -it curl sh 

cd /work/public1


curl -H "Content-Type: application/json" \
  -XPUT "http://localhost:9200/index" \
  --data-binary "@mapping.json"

curl \
  -H "Content-Type: application/x-ndjson" \
  -XPOST "http://localhost:9200/index/_bulk" \
  --data-binary "@test.json"



# If you want to compile the code
#docker build -t my-golang .
#docker run -p 8080:8080 -e ELASTICSEARCH_URL="http://elasticsearch:9200" --network localhost my-golang

cd search-api
go run main.go


