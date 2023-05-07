package main

import (
	"http-server/routers"
	"http-server/task"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	go task.Start() // 启动异步任务worker
	// go task.StartCron() // 启动定时任务

	r := gin.Default()

	//加载静态文件和模板
	r.LoadHTMLGlob("templates/*") //加载模板（页面）

	//使用cookie来进行登录验证
	store := cookie.NewStore([]byte("golang"))   //golang 这是字符串的意思是密钥
	r.Use(sessions.Sessions("mysession", store)) //名称

	r.Static("/statics", "./statics") //加载静态文件 css js image 。。。
	// 路由分组
	routers.UserRouterInit(r)
	routers.DefaultRouters(r)

	r.Run(":8080")

}
