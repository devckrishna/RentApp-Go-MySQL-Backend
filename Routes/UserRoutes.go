package routes

import (
	controllers "github.com/devckrishna/rentapp/Controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(ctx *gin.Engine) {
	ctx.GET("/users", controllers.GetUsers)
	ctx.GET("users/:id", controllers.GetUserById)
	ctx.POST("/users/signup", controllers.SignUp)
	ctx.DELETE("/users/:id", controllers.DeleteUser)
	ctx.POST("/users/login", controllers.LoginUser)
	ctx.POST("/users/enablehost/:id", controllers.EnableHost)
	ctx.POST("/users/disablehost/:id", controllers.DisableHost)
}
