package main

import (
	"fmt"
	"go.uber.org/zap"
	"lab_sys/global"
	"lab_sys/initialize"
)

func main() {
	initialize.InitConfig()
	//初始化数据库
	initialize.InitDB()
	initialize.InitRedis()
	//初始化日志
	initialize.InitLogger()
	//初始化我们的路由服务
	r := initialize.Routers()
	//
	err := r.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("this is hello ", zap.String("user", "liuan"))
	}
}
