package main

import "lab_sys/lab_sys/routes"

func main() {
	route := routes.Router()

	route.StaticFile("lab_sys/fronted/login.html", "./fronted/index.html")
	route.Run(":8080")
}
