package initialize

import (
	"github.com/gin-gonic/gin"
	"lab_sys/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	v1Groups := r.Group("/v1")

	router.UserRouter(v1Groups)

	return r
}
