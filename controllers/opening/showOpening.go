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
	r.SendSuccess(ctx, "show opening", opening)
}
