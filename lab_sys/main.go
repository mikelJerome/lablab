package main

import (
	"lab_sys/lab_sys/initialize"
	"lab_sys/lab_sys/routes"
)

func main() {
	//初始化数据库
	initialize.InitDB()
	//初始化日志
	initialize.InitLogger()

	route := routes.Router()
	route.Run(":8080")
}
