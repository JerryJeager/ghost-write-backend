package cmd

import (
	"log"
	"os"

	"github.com/JerryJeager/ghost-write-backend/api"
	"github.com/JerryJeager/ghost-write-backend/manualwire"
	"github.com/JerryJeager/ghost-write-backend/middleware"
	"github.com/gin-gonic/gin"
)

var userController = manualwire.GetUserController()
var messageController = manualwire.GetMessageController()

func ExecuteApiRoutes() {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())


	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!",
		})
	})


	v1 := r.Group("/api/v1")
	v1.GET("/info/openapi.yaml", func(c *gin.Context) {
		c.String(200, api.OpenApiDocs())
	})

	users := v1.Group("/users")
	users.POST("/signup", userController.CreateUser)
	users.POST("/login", userController.CreateToken)
	users.GET("/:user-id", userController.GetUser)

	message := v1.Group("/message")
	message.POST("/:user-id", messageController.CreateMessage)

	anonMessages := v1.Group("/message")
	anonMessages.Use(middleware.JwtAuthMiddleware())
	anonMessages.GET("/:user-id", messageController.GetMessages)


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}