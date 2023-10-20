package api

import (
	"GoOrderPackProject/api/controllers"
)

var server = controllers.Server{}

func Run() {
	server.Initialize()
	server.Run(":8080")
}
