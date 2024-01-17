package main

import (
	"example.com/final_project_-_REST_API/db"
	"example.com/final_project_-_REST_API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
