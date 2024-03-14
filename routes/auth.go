// routes/auth_routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuuheiigarashi/go-rest-api/handlers"
)

func SetupAuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signin", handlers.SignIn)
		authGroup.POST("/signup", handlers.SignUp)
	}
}
