package main

import (
	"go-mysql-api/config"
	"go-mysql-api/controller"
	"go-mysql-api/dto"
	"go-mysql-api/helper"
	"go-mysql-api/usecase/auth/repoauth"
	"go-mysql-api/usecase/auth/serviceauth"
	"go-mysql-api/usecase/book/repoimplbook"
	"go-mysql-api/usecase/book/serviceimplbook"
	"go-mysql-api/usecase/user/repoimpl"
	"go-mysql-api/usecase/user/serviceimpl"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	// authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	authRepo := repoauth.NewAuthGormRepository(db)
	authService := serviceauth.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	userRepo := repoimpl.NewGormRepository(db)
	userService := serviceimpl.NewService(userRepo)
	userController := controller.NewUserController(userService)

	bookRepo := repoimplbook.NewGormRepositoryBook(db)
	bookService := serviceimplbook.NewServiceBook(bookRepo)
	bookControoller := controller.NewBookController(bookService)

	router := gin.Default()
	// router.POST("/register", authController.Register)
	router.POST("/api/v1/login", authController.Login)
	router.POST("/api/v1/cek-token", authController.CekToken)

	authRoutes := router.Group("api/auth")
	authRoutes.Use(authMiddleware(*authService, *userService))
	{

		authRoutes.GET("/user", userController.GetUser)
		authRoutes.GET("/user/:id", userController.FindByIdUser)
		authRoutes.POST("/user/create-new", userController.CreateUser)
		authRoutes.POST("/user/update/:id", userController.UpdateUser)
		authRoutes.DELETE("/user/delete/:id", userController.DeleteUser)
		// book
		authRoutes.POST("book/create-new", bookControoller.CreateBook)
		authRoutes.GET("books", bookControoller.GetBooks)

	}
	router.Run(":3000")
	// listen and serve on 0.0.0.0:8080
}

func authMiddleware(authService serviceauth.Service, userService serviceimpl.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claims["user_id"].(float64))

		user, err := userService.UserFindById(dto.GetUserByIDRequest{ID: uint64(userID)})
		c.JSON(http.StatusOK, gin.H{
			"message": "user all",
			"users":   user,
		})
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}

// token git
// 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
// https://www.youtube.com/watch?v=yGTMQ5e-T5E&list=PLkVx132FdJZlTc_1gucKZ00b_s45DQlVQ&index=4
