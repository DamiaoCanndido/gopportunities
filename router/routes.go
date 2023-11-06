package router

import (
	"github.com/DamiaoCanndido/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine){
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningsHandler)
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningsHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningsHandler)
	}

	r.Run(":3333")
}

func CreateOpeningHandler() {
	panic("unimplemented")
}