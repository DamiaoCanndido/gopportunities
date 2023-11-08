package handler

import (
	"fmt"
	"net/http"

	"github.com/DamiaoCanndido/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningsHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
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
		sendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error updating opening with id: %s", id))
	}
	sendSuccess(ctx, "update opening", opening)
}
