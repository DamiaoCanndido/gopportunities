package opening

import (
	"fmt"
	"net/http"

	h "github.com/DamiaoCanndido/gopportunities/config"
	"github.com/DamiaoCanndido/gopportunities/entities"
	r "github.com/DamiaoCanndido/gopportunities/helpers"
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
func DeleteOpeningController(ctx *gin.Context) {
	db := h.GetPostgreSQL()
	id := ctx.Params.ByName("id")
	if id == "" {
		r.SendError(
			ctx, http.StatusBadRequest, ErrParamIsRequired("id", "parameter").Error())
		return
	}
	opening := entities.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		r.SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	if err := db.Delete(&opening).Error; err != nil {
		r.SendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting opening with id: %s", id))
	}
	r.SendSuccess(ctx, "delete opening", opening)
}
