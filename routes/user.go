// routes/user.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuuheiigarashi/go-rest-api/handlers"
)

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", handlers.GetUsers)
		userGroup.GET("/:id", handlers.GetUser)
		userGroup.POST("/", handlers.CreateUser)
		userGroup.PUT("/:id", handlers.UpdateUser)
		userGroup.DELETE("/:id", handlers.DeleteUser)
	}
}
