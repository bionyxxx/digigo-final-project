package main

import (
	"Final_Project/configs"
	"Final_Project/controllers"
	"Final_Project/repositories"
	"Final_Project/routes"
	"Final_Project/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func InitApp() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = configs.InitDatabase()
	if err != nil {
		panic(err)
	}
}

func StartApp() {
	var router = gin.New()
	repo := repositories.NewRepo(configs.GORM.DB)
	service := services.NewService(repo)
	server := controllers.NewHttpServer(service)
	routes.ApiInit(router, server)
	var PORT = os.Getenv("APP_PORT")
	err := router.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		panic(err)
	}
}
