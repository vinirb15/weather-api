package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinirb15/go-api/schemas"
)

// @BasePath /api/v1

// @Summary Create weather
// @Description Create a new job weather
// @Tags weathers
// @Accept json
// @Produce json
// @Param request body CreateweatherRequest true "Request body"
// @Success 200 {object} CreateweatherResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /weather [post]
func CreateweatherHandler(ctx *gin.Context) {
	request := CreateweatherRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	weather := schemas.Weather{
		Location:    request.Location,
		Name:        request.Name,
		Temperature: request.Temperature,
	}

	if err := db.Create(&weather).Error; err != nil {
		logger.Errorf("error creating weather: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating weather on database")
		return
	}

	sendSuccess(ctx, "create-weather", weather)
}
