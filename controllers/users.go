package controllers

import (
	"log"
	"net/http"

	models "example.com/plan-my-event/models"
	"example.com/plan-my-event/utils"
	"github.com/gin-gonic/gin"
)

type IUsersController interface {
	Signup(*gin.Context)
	Login(*gin.Context)
}

// @Summary	Signup for an user
// @Description	Registers an user in the app
// @Tags	users
// @Param user body models.User true "user body"
// @Success	201 {object} map[string]any
// @Failure	400 {object} map[string]any
// @Failure	500 {object} map[string]any
// @Router /signup [post]
func (userCont controllerImpl) Signup(context *gin.Context) {
	var user models.User
	// bind payload with user model
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Println("Error in parsing the user : " + err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get the data from the client"})
		return
	}
	log.Println("Started Signing up the user")
	// Service call to sign up the user
	custErr := userCont.service.SignUp(context, &user)
	if custErr.Err != nil {
		log.Println("Error in signingup the user : " + custErr.Err.Error())
		context.JSON(custErr.StatusCode, gin.H{"message": custErr.Err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})
}

// @Summary	Login for an user
// @Description	Logins an user in the app
// @Tags	users
// @Param user body models.User true "user body"
// @Success	200 {object} map[string]any
// @Failure	400 {object} map[string]any
// @Failure	500 {object} map[string]any
// @Router /login [post]
func (userCont controllerImpl) Login(context *gin.Context) {
	var user models.User
	// bind payload to user model
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Println("Error in parsing the user : " + err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	log.Println("Started login in the user")
	// Service call to login the user
	custErr := userCont.service.Login(context, &user)
	if custErr.Err != nil {
		log.Println("Error in logging in the user : " + custErr.Err.Error())
		context.JSON(custErr.StatusCode, gin.H{"message": custErr.Err.Error()})
		return
	}
	log.Println("Started generation of jwt token")
	// Generate new token for authentication
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		log.Println("Error in generating the token for the user")
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login Successful !", "token": token})
}
