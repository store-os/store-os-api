package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/store-os/store-os-api/api"
	"github.com/store-os/store-os-api/httputil"
)

// Products godoc
// @Summary Products endpoint
// @Description List products
// @Tags search
// @Accept  json
// @Produce  json
// @Success 200 {array} api.Product
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /products [get]
func (c *Controller) ListProducts(ctx *gin.Context) {

	body, err := api.ListProducts()

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": body,
	})
}

// Products godoc
// @Summary Products endpoint
// @Description get product by ID
// @Tags search
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {array} api.Product
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /products/{id} [get]
func (c *Controller) OneProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	body, err := api.OneProduct(id)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, body)
}
