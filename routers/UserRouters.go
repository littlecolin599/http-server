package routers

import (
	"http-server/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRoutters := r.Group("/login")
	{
		userRoutters.GET("/", user.UserController{}.Index) // 访问 localhost:8080/login -> 触发 user.UserController{}.Index方法
		userRoutters.POST("/doLogin", user.UserController{}.DoLogin)
		userRoutters.GET("/welcome", user.UserController{}.Welcome)
		userRoutters.POST("/5", user.UserController{}.GetImage)
		userRoutters.POST("/6", user.UserController{}.GetVideo)
		userRoutters.POST("/7", user.UserController{}.DoSomething)
	}
}
