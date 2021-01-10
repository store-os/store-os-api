package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/store-os/store-os-api/api"
	"github.com/store-os/store-os-api/httputil"
)

// Autocomplete godoc
// @Summary List autocomplete
// @Description get autocomplete
// @Tags autocomplete
// @Accept  json
// @Produce  json
// @Param client path string true "client"
// @Param q query string false "name search by q" Format(q)
// @Success 200 {array} api.Autocomplete
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /{client}/autocomplete [get]
func (c *Controller) Autocomplete(ctx *gin.Context) {

	client := ctx.Param("client")
	if client == "" {
		log.Println("Client no-specified")
		return
	}

	q := ctx.Request.URL.Query().Get("q")

	category, _ := ctx.Request.URL.Query()["category"]
	subcategory := ctx.Request.URL.Query()["subcategory"]
	subsubcategory := ctx.Request.URL.Query()["subsubcategory"]
	from := ctx.Request.URL.Query().Get("from")
	to := ctx.Request.URL.Query().Get("to")

	body, err := api.SearchAutocomplete(client, q, category, subcategory, subsubcategory, from, to)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"autocomplete": body,
	})
}
