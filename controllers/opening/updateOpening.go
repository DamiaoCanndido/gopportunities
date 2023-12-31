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

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body UpdateOpeningRequest true "Request body"
// @Param id path int true "Opening ID"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{id} [put]
func UpdateOpeningController(ctx *gin.Context) {
	logger := h.GetLogger("logger")
	db := h.GetPostgreSQL()
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		r.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
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
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if err := db.Save(&opening).Error; err != nil {
		r.SendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error updating opening with id: %s", id))
	}
	r.SendSuccess(ctx, "update opening", opening)
}
