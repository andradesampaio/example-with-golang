package routes

import "github.com/gin-gonic/gin"

func ClientRoutes(route *gin.Engine) *gin.Engine{

	userRouter := route.Group("/v1") 
	{
		userRouter.GET("/client", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "Get all client",
			})
		})

		userRouter.POST("/client/:id", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "Client user by id",
			})
		})
	}

	return route
}