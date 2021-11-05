package main

import (
	"github.com/gin-gonic/gin"
	"web/poject/route"
)

func main() {
	server := gin.Default()
	route.NewRoute(server).AddStudentRoute()
	server.Run(":8080")
}
