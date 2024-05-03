package main

import (
	"fmt"
	"group-routes/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")
	router := gin.Default()
	routes.ClientRoutes(router)
	routes.UserRoutes(router)

	router.Use(cors.Default())
	router.Run(":8080")
}