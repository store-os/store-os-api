# store-os-api
API based in Go

## Usage

### Keep it simple

For those who wants keep it simple and avoid install a lot of things:

```sh
docker-compose up -d

```
### Windows
```sh
docker run -v "C:/Hugo/themes/store-os/store-os-api/public/curl.sh:/work/curl.sh" -v "C:/Hugo/themes/store-os/store-os-api/public/test_search.json:/work/test_search.json" -v "C:/Hugo/themes/store-os/store-os-api/public/test_blog.json:/work/test_blog.json" -v "C:/Hugo/themes/store-os/store-os-api/public/search_template.json:/work/search_template.json" -it "ellerbrock/alpine-bash-curl-ssl" sh
```

### Start using it
1. Add comments to your API source code, [See Declarative Comments Format](https://swaggo.github.io/swaggo.io/declarative_comments_format/).
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:
```sh
$ go get -u github.com/swaggo/swag/cmd/swag
```

3. Run the [Swag](https://github.com/swaggo/swag) in your Go project root folder which contains `main.go` file, [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`).
```sh
$ swag init
```
4. Run Store OS API 
```sh 
$ go run main.go
```

## Docker 
```sh
brew install docker
docker network create localhost
```

## Elasticsearch

```sh
docker run -d --network localhost --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.1
```

Elasticsearch will be located in localhost:9200

## Cerebro 

To monitor elasticsearch indexes

```sh
docker run -d --network localhost -p 9000:9000 --name cerebro yannart/cerebro:latest
```
Hit localhost:9000 in the browser to see the interface and enter http://elasticsearch:9200 to reach Elasticsearch

## Kibana

To enter index mapping and to explore new queries/options in Elasticsearch

```sh
docker run --link 14833cb3405f:elasticsearch -p 5601:5601 --network localhost docker.elastic.co/kibana/kibana:7.6.1
```
Hit localhost:5601 in the browser to enter Kibana


## API 

```sh
docker build -t my-golang .
docker run -p 8080:8080 -e ELASTICSEARCH_URL="http://elasticsearch:9200" --network localhost my-golang
docker logs -f container_id
```


## Site Example 

1. Index Template
```sh
curl \
  -H "Content-Type: application/json" \
  -XPUT "http://localhost:9200/_template/search_template" \
  --data-binary "@public/search_template.json"
```
or using Kibana to insert the index mapping 

2. Index search data

```sh
 curl \
  -H "Content-Type: application/x-ndjson" \
  -XPOST "http://localhost:9200/index_search/_bulk" \
  --data-binary "@public/test_search.json"
```
3. Hit API 
```sh 
http://localhost:8080/search?q=home
http://localhost:8080/suggest?q=hom
http://localhost:8080/health
```

### Autocomplete 

[Search as you type](https://www.elastic.co/guide/en/elasticsearch/reference/7.9/search-as-you-type.html)


### Endpoints

[Gin Tonic](https://github.com/gin-gonic/gin)

### Swagger

[Gin Swagger](https://github.com/swaggo/gin-swagger)