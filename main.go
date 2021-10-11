package main

import (
	"go-mysql-api/config"
	"go-mysql-api/controller"
	"go-mysql-api/usecase/user/repoimpl"
	"go-mysql-api/usecase/user/serviceimpl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	userRepo := repoimpl.NewGormRepository(db)
	userService := serviceimpl.NewService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.GET("/user", userController.GetUser)
		authRoutes.GET("/user/:id", userController.FindByIdUser)
		authRoutes.POST("/user/create-new", userController.CreateUser)

	}
	router.Run(":3000")
	// listen and serve on 0.0.0.0:8080
}

// token git
// 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
// https://www.youtube.com/watch?v=yGTMQ5e-T5E&list=PLkVx132FdJZlTc_1gucKZ00b_s45DQlVQ&index=4
