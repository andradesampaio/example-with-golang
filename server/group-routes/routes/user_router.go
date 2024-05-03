package routes

import "github.com/gin-gonic/gin"

func UserRoutes(route *gin.Engine) *gin.Engine{

	userRouter := route.Group("/v1")
	{
		userRouter.GET("/user", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "Get all users",
			})
		})
		userRouter.POST("/user/:id", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message": "Get user by id",
			})
		})
	}
	return route
}

