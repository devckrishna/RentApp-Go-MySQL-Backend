package routes

import (
	controllers "github.com/devckrishna/rentapp/Controllers"
	"github.com/gin-gonic/gin"
)

func RequestRoutes(ctx *gin.Engine) {
	ctx.POST("/request/create", controllers.CreateRequest)
	ctx.GET("/request/:id", controllers.GetRequestByPropertyId)
	ctx.GET("/request/myrequests/:id", controllers.GetAllRequestsByUser)
}
