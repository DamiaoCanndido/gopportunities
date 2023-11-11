package router

import (
	c "github.com/DamiaoCanndido/gopportunities/controllers/opening"
	"github.com/gin-gonic/gin"
)

func openingRoutes(r *gin.Engine, basePath string) {
	op := r.Group(basePath)
	{
		op.GET("/opening", c.ShowOpeningsController)
		op.GET("/opening/:id", c.ShowOpeningController)
		op.POST("/opening", c.CreateOpeningController)
		op.PUT("/opening/:id", c.UpdateOpeningController)
		op.DELETE("/opening/:id", c.DeleteOpeningController)
	}
}
