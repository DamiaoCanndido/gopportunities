package opening

import (
	"net/http"

	h "github.com/DamiaoCanndido/gopportunities/config"
	"github.com/DamiaoCanndido/gopportunities/entities"
	r "github.com/DamiaoCanndido/gopportunities/helpers"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningController(ctx *gin.Context) {
	logger := h.GetLogger("logger")
	db := h.GetPostgreSQL()
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		r.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := entities.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		r.SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}
	r.SendSuccess(ctx, "create opening", opening)
}
