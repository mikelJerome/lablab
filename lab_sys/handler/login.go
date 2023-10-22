package handler

import (
	"github.com/gin-gonic/gin"
	"lab_sys/lab_sys/utils"
)

// 设置一个登录请求

func LoginRegister(c *gin.Context) {
	// 获取请求的参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	// 验证码
	verifycode := c.PostForm("verifycode")
	emaildata, exist := c.Get("emaildata")
	if !exist {
		c.JSON(400, gin.H{
			"error": "未找到有效的email值",
		})
		return
	}
	// !!!!!!!!!!!!!!反复观看!!!!!!!!!!!!!!!!!!!!
	email := emaildata.(*EmailData).Email
	verifycoderedis := utils.Redis.GetCache(email)

}
