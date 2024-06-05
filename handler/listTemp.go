package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinirb15/go-api/schemas"
)

// @BasePath /api/v1

// @Summary List weathers
// @Description List all job weathers
// @Tags weathers
// @Accept json
// @Produce json
// @Success 200 {object} ListweathersResponse
// @Failure 500 {object} ErrorResponse
// @Router /weathers [get]
func ListweathersHandler(ctx *gin.Context) {
	weathers := []schemas.Weather{}

	if err := db.Find(&weathers).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing weathers")
		return
	}

	sendSuccess(ctx, "list-weathers", weathers)
}
