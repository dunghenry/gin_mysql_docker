package routes

import (
	"gin/mysql/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.GET("/", controllers.GetUsers)
	users.POST("/", controllers.AddUser)
	users.GET("/:id", controllers.GetUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
}
