package initialize

import (
	"github.com/spf13/viper"
	"lab_sys/config"
	"lab_sys/global"
)

func InitConfig() {
	v := viper.New()

	// 读取文件路径
	v.SetConfigFile("./setting.yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	global.Settings = serverConfig

}
