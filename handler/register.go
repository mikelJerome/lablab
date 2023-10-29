package handler

import (
	"github.com/gin-gonic/gin"
	"lab_sys/lab_sys/database"
	"lab_sys/lab_sys/model"
	"lab_sys/lab_sys/utils"
)

// 设置一个登录请求

func Register(c *gin.Context) {
	user := model.User{}
	// 获取请求的参数

	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	//验证码
	verifycode := c.PostForm("verifycode")

	//email := emailcode

	verifycoderedis := utils.Redis.SGet(user.Email)

	if verifycoderedis == "" {
		c.JSON(400, gin.H{
			"error": "验证码错误或者没有填写",
		})
		return
	}

	if verifycode == verifycoderedis {
		_, err := database.FindPhone(user.Phone)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "用户已经存在",
			})
			return
		}
		user1, err := database.CreateUser(user)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "创建失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": user1,
		})

	}
}
