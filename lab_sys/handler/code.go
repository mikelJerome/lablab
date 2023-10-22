package handler

import (
	"github.com/gin-gonic/gin"
	"lab_sys/lab_sys/utils"
	"time"
)

type EmailData struct {
	Email string
}

// 获取验证码请求的api
func RequestVerifyCode(c *gin.Context) {
	email := c.PostForm("eamil")
	verifycode := utils.Email(email)
	_ = utils.NewRedisClient()
	utils.Redis.Set(email, verifycode, time.Minute)
	emaildata := &EmailData{
		Email: email,
	}
	c.Set("emaildata", emaildata)
	c.JSON(200, gin.H{
		"message": "验证码已发送",
	})

}
