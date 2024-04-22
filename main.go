package main

import (
	"github.com/gin-gonic/gin"
	"rest.api/db"
	"rest.api/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
