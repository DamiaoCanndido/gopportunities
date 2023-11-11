package router

import (
	c "github.com/DamiaoCanndido/gopportunities/controllers"
	docs "github.com/DamiaoCanndido/gopportunities/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	c.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	op := r.Group(basePath)
	{
		op.GET("/opening", c.ShowOpeningsController)
		op.GET("/opening/:id", c.ShowOpeningController)
		op.POST("/opening", c.CreateOpeningController)
		op.PUT("/opening/:id", c.UpdateOpeningController)
		op.DELETE("/opening/:id", c.DeleteOpeningController)
	}
	auth := r.Group(basePath)
	{
		auth.GET("/auth/login/google", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"login": "google",
			})
		})
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":3333")
}
