package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"lab_sys/database"
	"lab_sys/response"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	phone := c.Query("phone")
	user, err := database.FindUserByPhone(phone)
	if err != nil {
		// 电话号码不匹配
		zap.L().Error("您正在尝试没有注册的电话号码")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid phone number"})
		return
	}
	if user.Username != username {
		// 用户不存在
		zap.L().Error("用户不存在并且没有查到")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}
	response.Success(c, http.StatusOK, 200, "success")
}
