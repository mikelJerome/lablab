package main

import "lab_sys/lab_sys/routes"

func main() {
	route := routes.Router()
	route.Run(":8080")
}
