package main

import (
	"back/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}
