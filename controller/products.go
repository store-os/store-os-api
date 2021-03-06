package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/store-os/store-os-api/api"
	"github.com/store-os/store-os-api/httputil"
)

// Products godoc
// @Summary Products endpoint
// @Description List products
// @Tags products
// @Accept  json
// @Produce  json
// @Param client path string true "client"
// @Param q query string false "name search by q" Format(q)
// @Param page query string false "paging number" Format(page=1)
// @Param category query string false "category filter" Format(category=)
// @Param subcategory query string false "subcategory filter" Format(subcategory=)
// @Param subsubcategory query string false "subsubcategory filter" Format(subsubcategory=)
// @Param from query int false "from price" Format(from)
// @Param to query int false "to price" Format(to)
// @Param page query int false "page number" Format(page)
// @Param fieldsort query string false "fieldsort final_price or title.keyword" Format(fieldsort)
// @Param order query string false "order (asc or desc)" Format(order)
// @Param size query string false "size" Format(size)
// @Success 200 {array} api.Product
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /{client}/products [get]
func (c *Controller) Search(ctx *gin.Context) {

	client := ctx.Param("client")
	if client == "" {
		log.Println("Client no-specified")
		return
	}
	q := ctx.Request.URL.Query().Get("q")
	category, _ := ctx.Request.URL.Query()["category"]
	subcategory := ctx.Request.URL.Query()["subcategory"]
	subsubcategory := ctx.Request.URL.Query()["subsubcategory"]
	fieldSort := ctx.Request.URL.Query().Get("fieldsort")
	order := ctx.Request.URL.Query().Get("order")
	from := ctx.Request.URL.Query().Get("from")
	to := ctx.Request.URL.Query().Get("to")
	page := ctx.Request.URL.Query().Get("page")
	size := ctx.Request.URL.Query().Get("size")

	if page == "" {
		page = "0"
	}
	if fieldSort == "" {
		fieldSort = "title.keyword"
	}
	if order == "" {
		order = "asc"
	}
	if size == "" {
		size = "20"
	}
	var body api.SearchResponse
	var err error
	//log.Println(category)
	body, err = api.Search(client, q, page, category, subcategory, subsubcategory, fieldSort, order, from, to, size)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, body)
}

// Products godoc
// @Summary OneProduct endpoint
// @Description get product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param client path string true "client"
// @Param id path int true "Product ID"
// @Success 200 {object} api.OneProductResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /{client}/products/{id} [get]
func (c *Controller) OneProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	client := ctx.Param("client")
	if client == "" {
		log.Println("Client no-specified")
		return
	}

	body, err := api.OneProduct(client, id)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, body)
}

// Products godoc
// @Summary Aggs endpoint
// @Description get aggs
// @Tags products
// @Accept  json
// @Produce  json
// @Param client path string true "client"
// @Success 200 {object} api.AggsResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /{client}/aggs [get]
func (c *Controller) Aggs(ctx *gin.Context) {
	client := ctx.Param("client")
	if client == "" {
		log.Println("Client no-specified")
		return
	}

	body, err := api.Aggs(client)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, body)
}
