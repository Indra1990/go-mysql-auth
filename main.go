package main

import (
	"go-mysql-api/config"
	"go-mysql-api/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := gin.Default()
	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)

	}
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// token git
// 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
// https://www.youtube.com/watch?v=yGTMQ5e-T5E&list=PLkVx132FdJZlTc_1gucKZ00b_s45DQlVQ&index=4
