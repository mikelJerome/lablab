package handler

import (
	"lab_sys/utils"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

//	type EmailData struct {
//		Email string
//	}
var (
	emailMutex sync.Mutex
	emailcode  string
)

// 获取验证码请求的api
func RequestVerifyCode(c *gin.Context) {
	email := c.PostForm("email")
	verifycode := utils.Email(email)

	// 获取redis链接

	redis, err := utils.GetRedisClient()
	if err != nil {
		panic(err)
	}

	// redis存进去
	redis.SSet(email, verifycode)

	emailMutex.Lock()
	emailcode = email
	emailMutex.Unlock()

	//c.Set("emaildata", emaildata)
	c.JSON(http.StatusOK, gin.H{
		"message": "验证码已发送",
		"code":    verifycode,
	})

}
