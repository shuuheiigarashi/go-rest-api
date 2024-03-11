// main.go
package main

import (
	"go-rest-api/docs"
	"go-rest-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Swaggerのエンドポイントを設定
	docs.SwaggerInfo.Title = "Your API"
	docs.SwaggerInfo.Version = "1.0"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ここにAPIルートを追加
	r.GET("/ping", handlers.Ping)

	r.Run(":8080")
}
