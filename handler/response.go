package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinirb15/go-api/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateWeatherResponse struct {
	Message string                  `json:"message"`
	Data    schemas.WeatherResponse `json:"data"`
}

type DeleteWeatherResponse struct {
	Message string                  `json:"message"`
	Data    schemas.WeatherResponse `json:"data"`
}
type ShowWeatherResponse struct {
	Message string                  `json:"message"`
	Data    schemas.WeatherResponse `json:"data"`
}
type ListweathersResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.WeatherResponse `json:"data"`
}
type UpdateWeatherResponse struct {
	Message string                  `json:"message"`
	Data    schemas.WeatherResponse `json:"data"`
}
