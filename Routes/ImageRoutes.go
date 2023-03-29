package routes

import (
	controllers "github.com/devckrishna/rentapp/Controllers"
	"github.com/gin-gonic/gin"
)

func ImageRoutes(ctx *gin.Engine) {
	ctx.POST("/image/add", controllers.AddImage)
	ctx.GET("/image/getpropertyimages/:id", controllers.GetPropertyImages)
}
