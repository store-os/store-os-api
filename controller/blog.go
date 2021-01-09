package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/store-os/store-os-api/api"
	"github.com/store-os/store-os-api/httputil"
)

// Products godoc
// @Summary Blog endpoint
// @Description List posts
// @Tags posts
// @Accept  json
// @Produce  json
// @Param client path string true "client"
// @Param page query string false "paging number" Format(page=1)
// @Success 200 {array} api.Product
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /{client}/blog [get]
func (c *Controller) ListPosts(ctx *gin.Context) {
	page := ctx.Request.URL.Query().Get("page")
	size := ctx.Request.URL.Query().Get("size")

	if page == "" {
		page = "0"
	}

	if size == "" {
		size = "20"
	}

	client := ctx.Param("client")
	if client == "" {
		log.Println("Client no-specified")
		return
	}

	body, hits, err := api.ListPosts(client, page, size)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"posts": body,
		"hits":  hits,
	})
}

// Products godoc
// @Summary Blog endpoint
// @Description get post by ID
// @Tags post
// @Accept  json
// @Produce  json
// @Param client path string true "client"
// @Param id path int true "Post ID"
// @Success 200 {array} api.Post
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /{client}/blog/{id} [get]
func (c *Controller) OnePost(ctx *gin.Context) {
	id := ctx.Param("id")
	client := ctx.Param("client")
	if client == "" {
		log.Println("Client no-specified")
		return
	}
	body, err := api.OnePost(client, id)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, body)
}
