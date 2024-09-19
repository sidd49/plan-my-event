package main

import (
	"log"
	"os"

	"example.com/plan-my-event/controllers"
	_ "example.com/plan-my-event/docs"
	"example.com/plan-my-event/repository"
	"example.com/plan-my-event/routes"
	"example.com/plan-my-event/service"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

//@title	Plan My Event API
//@version	1.0
//@description	This is a event management application

// @contact.name Siddhant
// @contact.email siddhant.sonar@hp.com
// @host	localhost:8080
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("No Env File Found")
		return
	}
	// db := repository.New()
	// db.Init()

	mongoDb := repository.NewMongo()
	mongoDb.InitMongo()
	service := service.NewService(mongoDb)
	controller := controllers.NewController(service)

	server := gin.Default()

	routes.RegisterRoutes(server, controller)

	server.Run(":" + os.Getenv("PME_PORT"))
}
