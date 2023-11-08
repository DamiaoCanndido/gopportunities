package handler

import (
	"fmt"
	"net/http"

	"github.com/DamiaoCanndido/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening ID"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{id} [delete]
func DeleteOpeningsHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		sendError(
			ctx, http.StatusBadRequest, errParamIsRequired("id", "parameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting opening with id: %s", id))
	}
	sendSuccess(ctx, "delete opening", opening)
}
