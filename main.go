package main

import (
	"lab_sys/initialize"
	"lab_sys/routes"
)

func main() {
	//初始化数据库
	initialize.InitDB()
	//初始化日志
	initialize.InitLogger()

	route := routes.Router()

	err := route.Run(":8000")
	if err != nil {
		println(err)
	}
}
