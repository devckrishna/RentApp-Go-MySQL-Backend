package routes

import (
	controllers "github.com/devckrishna/rentapp/Controllers"
	"github.com/gin-gonic/gin"
)

func SavedRoutes(ctx *gin.Engine) {
	ctx.POST("/saved/add", controllers.AddToSaved)
	ctx.GET("/saved/:id", controllers.GetSaved)
}
