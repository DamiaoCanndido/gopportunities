package opening

import (
	"net/http"

	h "github.com/DamiaoCanndido/gopportunities/controllers"
	"github.com/DamiaoCanndido/gopportunities/entities"
	r "github.com/DamiaoCanndido/gopportunities/helpers"
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
	_, db := h.InitializeHandler()
	openings := []entities.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		r.SendError(ctx, http.StatusInternalServerError, "error show openings")
		return
	}
	r.SendSuccess(ctx, "list openings", openings)
}
