package handler

import (
	"fmt"
	"net/http"

	"github.com/DamiaoCanndido/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningsHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		sendError(
			ctx, http.StatusBadRequest, errParamIsRequired("id", "quueryParameter").Error())
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
