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
// @Description post product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param client path string true "client"
// @Param id path int true "Product ID"
// @Success 200 {array} api.Product
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /{client}/products/{id} [post]
func (c *Controller) EditProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	client := ctx.Param("client")
	if client == "" {
		log.Println("Client no-specified")
		return
	}
	
	var product api.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		return
	}
	body, err := api.EditProduct(client, id, product)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, body)
}