package helpers

import (
	"fmt"
	"net/http"

	"github.com/DamiaoCanndido/gopportunities/entities"
	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                   `json:"message"`
	Data    entities.OpeningResponse `json:"data"`
}

type ShowOpeningsResponse struct {
	Message string                     `json:"message"`
	Data    []entities.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string                   `json:"message"`
	Data    entities.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string                   `json:"message"`
	Data    entities.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                   `json:"message"`
	Data    entities.OpeningResponse `json:"data"`
}
