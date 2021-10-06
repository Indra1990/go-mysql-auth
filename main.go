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
