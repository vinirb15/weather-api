package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinirb15/go-api/schemas"
)

// @BasePath /api/v1

// @Summary Update weather
// @Description Update a job weather
// @Tags weathers
// @Accept json
// @Produce json
// @Param id query string true "weather Identification"
// @Param weather body UpdateweatherRequest true "weather data to Update"
// @Success 200 {object} UpdateweatherResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /weather [put]
func UpdateweatherHandler(ctx *gin.Context) {
	request := UpdateweatherRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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
	// Update weather

	if request.Name != "" {
		weather.Name = request.Name
	}

	if request.Temperature > -100 {
		weather.Temperature = request.Temperature
	}

	if request.Location != "" {
		weather.Location = request.Location
	}

	// Save weather
	if err := db.Save(&weather).Error; err != nil {
		logger.Errorf("error updating weather: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating weather")
		return
	}
	sendSuccess(ctx, "update-weather", weather)
}
