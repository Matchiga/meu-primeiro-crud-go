package main

import (
	"log"

	"github.com/Matchiga/meu-primeiro-crud-go/src/controller"
	"github.com/Matchiga/meu-primeiro-crud-go/src/controller/routes"
	"github.com/Matchiga/meu-primeiro-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init dependencies
	service := service.NewUserDomainService()
	userControler := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userControler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
