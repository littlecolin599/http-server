package main

import (
	"http-server/routers"
	"http-server/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	go task.Start() // 启动异步任务worker
	// go task.StartCron() // 启动定时任务

	r := gin.Default()

	//加载静态文件和模板
	r.LoadHTMLGlob("templates/*") //加载模板（页面）

	r.Static("/statics", "./statics") //加载静态文件 css js image 。。。
	r.GET("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})

	// 路由分组
	routers.UserRouterInit(r)
	routers.DefaultRouters(r)

	r.Run()

}
