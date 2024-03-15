// routes/todo.go

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuuheiigarashi/go-rest-api/handlers"
)

// SetupToDoRoutes はToDo関連のエンドポイントを設定します
func SetupToDoRoutes(r *gin.Engine) {
	todoGroup := r.Group("/todos")
	{
		todoGroup.GET("/", handlers.GetToDoList)
		todoGroup.POST("/", handlers.CreateToDoItem)
		todoGroup.PUT("/:id", handlers.UpdateToDoItem)
		todoGroup.DELETE("/:id", handlers.DeleteToDoItem)
	}
}
