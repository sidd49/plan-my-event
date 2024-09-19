package routes

import (
	"example.com/plan-my-event/commondef"
	"example.com/plan-my-event/controllers"
	"example.com/plan-my-event/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// function to route different apis with their controllers
func RegisterRoutes(server *gin.Engine, controller controllers.Controller) {

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pmeAPI := server.Group(commondef.BasePMEUrl + commondef.VersionPME)

	pmeAPI.GET("/events", controller.GetEvents)
	pmeAPI.GET("/events/:id", controller.GetEvent)

	pmeAPI.POST("/signup", controller.Signup)
	pmeAPI.POST("/login", controller.Login)

	authenticated := pmeAPI.Group("/")

	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", controller.CreateEvent)
	authenticated.PUT("/events/:id", controller.UpdateEvent)
	authenticated.DELETE("events/:id", controller.DeleteEvent)

	authenticated.POST("/events/:id/register", controller.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", controller.CancelRegistration)

}
