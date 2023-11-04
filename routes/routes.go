package routes

import (
	"lab_sys/handler"

	"github.com/gin-gonic/gin"
)

// 请求

func Router() *gin.Engine {
	r := gin.Default()

	//就在html模板
	r.LoadHTMLGlob("./frontend/*")
	r.GET("/rhtml", func(c *gin.Context) {
		c.HTML(200, "Register.html", nil)
	})
	r.GET("/reshtml", func(c *gin.Context) {
		c.HTML(200, "reserve.html", nil)
	})

	// post处理请求
	r.POST("/requestCode", handler.RequestVerifyCode)
	r.POST("/register", handler.Register)
	r.POST("/reserve", handler.Reserve)
	//<------------------------------------------------------------------------------>

	// get请求
	r.GET("/login", handler.Login)
	return r
}
