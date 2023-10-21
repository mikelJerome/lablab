package handler

import (
	"github.com/gin-gonic/gin"
	"lab_sys/lab_sys/model"
	"lab_sys/lab_sys/utils"
	"net/http"
	"time"
)

// 设置一个登录请求

func Login(c *gin.Context) {
	// 获取请求的参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	// 验证码
	verifycode := c.PostForm("verifycode")
	redisClient, _ := utils.NewRedisClient()
	verifycodeemail := utils.Email()
	redisClient.SeCachet(username, verifycodeemail, time.Duration(time.Now().Add(1*time.Second).Unix()))
	verifycoderedis, _ := redisClient.GetCache(username)
	if verifycode == verifycoderedis {
		c.JSON(http.StatusOK, gin.H{
			"userinfo": model.User{
				Username: username,
				Password: password,
				Phone:    phone,
				Email:    email,
			},
		})
	}
	if verifycode != verifycoderedis {
		c.JSON(404, gin.H{"error": "Invalid verification code"})
		return
	}
	if verifycode == "" {
		c.JSON(404, gin.H{"error": "Verification code expired"})
	}

}
