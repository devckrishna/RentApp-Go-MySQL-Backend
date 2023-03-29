package routes

import (
	controllers "github.com/devckrishna/rentapp/Controllers"
	"github.com/gin-gonic/gin"
)

func ReviewRoutes(ctx *gin.Engine) {
	ctx.POST("/review/create", controllers.CreateReview)
	ctx.GET("/review/:id", controllers.GetReviewByPropertyId)
}
