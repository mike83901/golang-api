package main

import (
	_ "GoApi/docs"
	"GoApi/handlers"
	"GoApi/middleware"
	"GoApi/repository"
	"GoApi/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GoApi Example API
// @version 1.0
// @description This is a sample server for GoApi.
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	addressBookRepo := repository.NewAddressBookRepo()
	addressBookService := service.NewAddressBookService(addressBookRepo)
	handler := handlers.NewHandler(addressBookService)

	router := gin.Default()

	router.GET("/", handler.HomePage)
	router.GET("/getToken", handler.GenerateTokenHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Protected routes
	protected := router.Group("/", middleware.AuthMiddleware())
	protected.GET("/getAddress", handler.GetAddressBookAll)
	protected.GET("/GetAddressBookByUserName", handler.GetAddressBookByUserName)

	router.Run(":8080")
}
