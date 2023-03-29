package routes

import (
	controllers "github.com/devckrishna/rentapp/Controllers"
	"github.com/gin-gonic/gin"
)

func PropertyRoutes(ctx *gin.Engine) {
	ctx.GET("/property", controllers.GetAllProperties)
	ctx.POST("/property", controllers.CreateProperty)
	ctx.GET("/property/:id", controllers.GetPropertyById)
	ctx.GET("/property/toprated", controllers.GetTopRatedProperties)
	ctx.GET("/property/nearyou/:city", controllers.GetPropertiesNearYou)
}
