package handler

import (
	"net/http"

	"github.com/DamiaoCanndido/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error show openings")
		return
	}
	sendSuccess(ctx, "list openings", openings)
}
