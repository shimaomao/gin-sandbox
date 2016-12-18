package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/controllers"
	"github.com/tsuyoshiwada/gin-sandbox/middleware"
)

func buildRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Auth
	authController := controllers.NewAuthController(db)
	auth := router.Group("/auth")
	{
		auth.POST("/", authController.Auth)
	}

	// Users
	userController := controllers.NewUserController(db)
	users := router.Group("/users")
	{
		users.Use(middleware.JWTMiddleware(db))
		users.GET("/", userController.GetAll)
	}

	return router
}