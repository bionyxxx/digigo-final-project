package main

import (
	"Final_Project/configs"
	"Final_Project/controllers"
	"Final_Project/repositories"
	"Final_Project/routes"
	"Final_Project/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

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
