package helpers

import (
	"fmt"
	"net/http"

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
