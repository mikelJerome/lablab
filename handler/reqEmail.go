package handler

import (
	"github.com/gin-gonic/gin"
	"lab_sys/global"
	"lab_sys/utils"
	"net/http"
	"strings"
	"time"
)

// 获取验证码请求的api
func RequestVerifyCode(c *gin.Context) {
	//var user model.User
	Email := c.PostForm("email")
	mailConf := utils.MailboxConf{
		Title:         global.DefaultMailTitle,
		RecipientList: Email,
		Sender:        global.DefaultSender,

		SPassword: global.DefaultSPassword,
		SMTPAddr:  global.DefaultSMTPAddr,
		SMTPPort:  global.DefaultSMTPPort,
	}
	global.Redis.Set(c, "Email"+" "+Email, strings.TrimSpace(mailConf.Email()), 5*time.Minute)
	c.JSON(http.StatusOK, gin.H{
		"massage": "验证码发送成功",
		"error":   "no error",
	})
}
