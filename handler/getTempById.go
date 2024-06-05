package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinirb15/go-api/schemas"
)

// @BasePath /api/v1

// @Summary Show weather
// @Description Show a job weather
// @Tags weathers
// @Accept json
// @Produce json
// @Param id query string true "weather identification"
// @Success 200 {object} ShowweatherResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /weather [get]
func ShowweatherHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	weather := schemas.Weather{}
	if err := db.First(&weather, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "weather not found")
		return
	}

	sendSuccess(ctx, "show-weather", weather)
}
