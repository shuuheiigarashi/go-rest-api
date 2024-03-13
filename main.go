// main.go
package main

import (
	"github.com/shuuheiigarashi/go-rest-api/docs"
	"github.com/shuuheiigarashi/go-rest-api/handlers"
	"github.com/shuuheiigarashi/go-rest-api/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Swaggerのエンドポイントを設定
	docs.SwaggerInfo.Title = "Your API"
	docs.SwaggerInfo.Version = "1.0"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ここにAPIルートを追加
	r.GET("/ping", handlers.Ping)
	routes.SetupUserRoutes(r)

	r.Run(":8080")
}
