package initialize

import (
	"go.uber.org/zap"
	"lab_sys/utils"
)

func InitRedis() {
	err := utils.NewRedisClient()
	if err != nil {
		zap.L().Info("this is redis 启动错误", zap.String("", ""))
	}
}
