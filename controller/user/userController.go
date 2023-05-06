package user

import (
	"encoding/json"
	"fmt"
	"http-server/models"
	"http-server/task"
	"net/http"

	"github.com/gin-contrib/sessions"
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
		userOfPhone := []models.User{}
		userOfEmail := []models.User{}
		//GORM 原生数据库操作进行封装  增删查改
		err1 := models.DB.Where("phone=?", userName).Find(&userOfPhone).Error
		err2 := models.DB.Where("email=?", userName).Find(&userOfEmail).Error
		if err1 != nil && err2 != nil {

		}

	}
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func (UserController) DoRegiseter(c *gin.Context) {
	register := c.PostForm("register")
	if register == "register" {
		email := c.PostForm("email")
		phone := c.PostForm("phone")
		password := c.PostForm("password")

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

			session := sessions.Default(c)            //开启session服务
			userinfo, _ := json.Marshal(user)         //序列化
			session.Set("userinfo", string(userinfo)) //string化  Set方法设置关键字
			err1 := session.Save()                    //Sava 保存
			if err1 == nil {
				c.Redirect(http.StatusMovedPermanently, "/")
			} else {
				c.JSON(http.StatusOK, "异常错误1")
			}
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
