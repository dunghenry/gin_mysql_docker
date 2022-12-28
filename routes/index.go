package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine) {
	getRoutes(r)
	str := fmt.Sprintf("Server running on http://localhost:%s", os.Getenv("PORT"))
	fmt.Println(str)
	r.Run(":" + os.Getenv("PORT"))
}

func getRoutes(r *gin.Engine) {
	api := r.Group("/api")
	UserRoutes(api)
}
