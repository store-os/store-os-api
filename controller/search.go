package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/store-os/store-os-api/api"
	"github.com/store-os/store-os-api/httputil"
)

// Search godoc
// @Summary Search endpoint
// @Description search query
// @Tags search
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q" Format(q)
// @Param page query string false "paging number" Format(page=1)
// @Success 200 {array} api.Product
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /search [get]
func (c *Controller) Search(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")
	page := ctx.Request.URL.Query().Get("page")

	if page == "" {
		page = "0"
	}
	body, hits, err := api.Search(q, page)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": body,
		"hits":     hits,
	})
}
