package controllers

import (
	"fmt"
	"net/http"

	"github.com/DamiaoCanndido/gopportunities/entities"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening ID"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{id} [get]
func ShowOpeningController(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		sendError(
			ctx, http.StatusBadRequest, errParamIsRequired("id", "parameter").Error())
		return
	}
	opening := entities.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	sendSuccess(ctx, "show opening", opening)
}
