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
	"go-mysql-api/usecase/languages/repolanguages"
	"go-mysql-api/usecase/languages/servicelanguages"
	"go-mysql-api/usecase/user/repoimpl"
	"go-mysql-api/usecase/user/serviceimpl"
	"go-mysql-api/usecase/userlanguage/repouserlanguage"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	authRepo := repoauth.NewAuthGormRepository(db)
	authService := serviceauth.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	userLanguageRepo := repouserlanguage.NewGormRepository(db)

	userRepo := repoimpl.NewGormRepository(db)
	userService := serviceimpl.NewService(userRepo, userLanguageRepo)
	userController := controller.NewUserController(userService)

	bookRepo := repoimplbook.NewGormRepositoryBook(db)
	bookService := serviceimplbook.NewServiceBook(bookRepo)
	bookControoller := controller.NewBookController(bookService)

	langRepo := repolanguages.NewGormRepositoryLanguages(db)
	langService := servicelanguages.NewServiceLanguages(langRepo)
	langController := controller.NewLanguagesController(*langService)

	router := gin.Default()
	router.POST("/api/v1/register", authController.Register)
	router.POST("/api/v1/login", authController.Login)
	router.POST("/api/v1/cek-token", authController.CekToken)

	authRoutes := router.Group("api/v1/auth")
	authRoutes.Use(authMiddleware(*authService, *userService))
	{
		authRoutes.GET("/user", userController.GetUser)
		authRoutes.GET("/user/:id", userController.FindByIdUser)
		authRoutes.POST("/user/create-new", userController.CreateUser)
		authRoutes.POST("/user/update/:id", userController.UpdateUser)
		authRoutes.DELETE("/user/delete/:id", userController.DeleteUser)
		// book
		authRoutes.GET("books", bookControoller.GetBooks)
		authRoutes.POST("books/create-new", bookControoller.CreateBook)
		authRoutes.POST("books/create-new-multiple", bookControoller.CreateBookMultiple)
		authRoutes.POST("books/update-multiple", bookControoller.UpdateBookMultiple)
		authRoutes.GET("books/:id", bookControoller.FindByIdBook)
		authRoutes.POST("books/:id/update", bookControoller.UpdateBook)
		authRoutes.DELETE("books/:id/delete", bookControoller.DeleteBook)
		authRoutes.DELETE("books/delete-multiple", bookControoller.DeleteBookMultiple)
		// language
		authRoutes.GET("lang/", langController.GetlistLang)
		authRoutes.POST("lang/create-new", langController.CreateLang)

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
