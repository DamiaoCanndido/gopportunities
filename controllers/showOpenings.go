package controllers

import (
	"net/http"

	"github.com/DamiaoCanndido/gopportunities/entities"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show openings
// @Description Show all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ShowOpeningsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningsController(ctx *gin.Context) {
	openings := []entities.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error show openings")
		return
	}
	sendSuccess(ctx, "list openings", openings)
}