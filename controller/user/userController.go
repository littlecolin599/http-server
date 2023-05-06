package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (UserController) DoLogin(c *gin.Context) {
	// TODO:登录校验
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func (UserController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{})
}
