package router

import (
	"github.com/gin-gonic/gin"
	"lab_sys/handler"
)

// 请求

func UserRouter(v1 *gin.RouterGroup) {
	UserRouter := v1.Group("/user")
	{
		UserRouter.GET("login", handler.Login)
		UserRouter.POST("reqCode", handler.RequestVerifyCode)
		UserRouter.POST("register", handler.Register)

	}

}
