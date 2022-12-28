package routes

import (
	"gin/mysql/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.GET("/", controllers.GetUsers)
}
