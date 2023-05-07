package user

import (
	"fmt"
	"http-server/models"
	"http-server/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (UserController) DoLogin(c *gin.Context) {
	// TODO:登录校验
	login := c.PostForm("login")
	if login == "login" {
		userName := c.PostForm("Username")
		password := models.Md5(c.PostForm("Password")) //123456
		fmt.Println("doLogin", userName, password)

		//实例化
		user := []models.User{}
		//GORM 原生数据库操作进行封装  增删查改
		result := models.DB.Where("phone=? AND password=?", userName, password).Find(&user)

		if result.RowsAffected == 0 {
			c.JSON(http.StatusInternalServerError, "用户未注册")
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		}
	}

}

func (UserController) DoRegiseter(c *gin.Context) {
	register := c.PostForm("register")
	if register == "register" {
		email := c.PostForm("email")
		phone := c.PostForm("phone")
		password := models.Md5(c.PostForm("password"))

		fmt.Println("doregister", register, email, phone)
		//判断格式

		//邮箱正则式
		// pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
		// reg := regexp.MustCompile(pattern)
		// if ok := reg.MatchString(email); !ok {
		// 	c.String(http.StatusOK, "邮箱格式不对")
		// 	c.Abort()
		// 	return
		// }
		//手机正则式
		// regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
		// reg = regexp.MustCompile(regular)
		// if ok := reg.MatchString(phone); !ok {
		// 	c.JSON(http.StatusOK, "手机号格式不正确")
		// 	return
		// }
		// password := models.Md5(c.PostForm("password"))

		user := models.User{
			Email:    email,
			Phone:    phone,
			Password: password,
		}

		err := models.DB.Save(&user).Error

		if err != nil {
			c.JSON(http.StatusOK, "注册失败")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

func (UserController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{})
}

func (UserController) GetImage(c *gin.Context) {
	c.HTML(http.StatusOK, "picture.html", gin.H{})
}

func (UserController) GetVideo(c *gin.Context) {
	c.HTML(http.StatusOK, "video.html", gin.H{})
}
func (UserController) DoSomething(c *gin.Context) {
	task.TaskAdd(1, 2)
}
