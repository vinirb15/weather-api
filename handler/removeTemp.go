package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinirb15/go-api/schemas"
)

// @BasePath /api/v1

// @Summary Delete weather
// @Description Delete a new job weather
// @Tags weathers
// @Accept json
// @Produce json
// @Param id query string true "weather identification"
// @Success 200 {object} DeleteweatherResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /weather [delete]
func DeleteweatherHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	weather := schemas.Weather{}
	// Find Weather
	if err := db.First(&weather, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("weather with id: %s not found", id))
		return
	}
	// Delete Weather
	if err := db.Delete(&weather).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting weather with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-weather", weather)
}
