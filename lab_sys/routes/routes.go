package routes

import (
	"github.com/gin-gonic/gin"
	"lab_sys/lab_sys/handler"
)

// 请求
func Router() *gin.Engine {
	r := gin.Default()
	// 处理请求
	r.POST("login", handler.Login)
	return r
}
