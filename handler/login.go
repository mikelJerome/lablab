package handler

import (
	"lab_sys/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	phone := c.Query("phone")
	user, err := database.FindUserByPhone(phone)
	if err != nil {
		// 电话号码不匹配
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid phone number"})
		return
	}
	if user.Username != username {
		// 用户不存在
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}
	c.JSON(200, gin.H{
		"message": "登录成功",
		"code":    user,
	})

	//fmt.Println
}
