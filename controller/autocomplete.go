package controller

import (
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
// @Param q query string false "name search by q" Format(q)
// @Success 200 {array} api.Suggestion
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /autocomplete [get]
func (c *Controller) Autocomplete(ctx *gin.Context) {

	q := ctx.Request.URL.Query().Get("q")

	body, err := api.Autocomplete(q)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"suggestions": body,
	})
}
