package routers

import (
	"http-server/controller/public"

	"github.com/gin-gonic/gin"
)

func DefaultRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", public.DefaultController{}.Index)
	}
}
