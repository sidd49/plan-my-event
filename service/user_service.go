package service

import (
	"errors"

	models "example.com/plan-my-event/models"
	"example.com/plan-my-event/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserService interface {
	SignUp(*gin.Context, *models.User) *utils.CustomError
	Login(*gin.Context, *models.User) *utils.CustomError
}

func (userService serviceImpl) SignUp(ctx *gin.Context, user *models.User) *utils.CustomError {
	check := utils.CheckValidEmail(user.Email)
	if !check {
		return utils.NewCustomError(400, errors.New("invalid email, please enter a valid email"))
	}
	if user.Password == "" {
		return utils.NewCustomError(400, errors.New("invalid password, please enter a valid password"))
	}
	if !utils.CheckValidPassword(user.Password) {
		return utils.NewCustomError(400, errors.New("invalid password, please enter a valid password"))
	}
	result := userService.Store.GetUserByEmail(user.Email)
	if result.Email != "" {
		return utils.NewCustomError(400, errors.New("user already present"))
	}

	userID, err := utils.ConvertIDtoString(primitive.NewObjectID().String())
	if err != nil {
		return utils.NewCustomError(500, err)
	}
	user.ID = userID
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.NewCustomError(500, err)
	}
	user.Password = hashedPassword
	err = userService.Store.SaveMongo(user)
	return utils.NewCustomError(500, err)
}

func (userService serviceImpl) Login(ctx *gin.Context, user *models.User) *utils.CustomError {
	check := utils.CheckValidEmail(user.Email)
	if !check {
		return utils.NewCustomError(400, errors.New("invalid email id, please enter a valid email"))
	}
	if user.Password == "" {
		return utils.NewCustomError(400, errors.New("password cannot be empty"))
	}
	err := userService.Store.ValidateMongoCredentials(user)
	return utils.NewCustomError(500, err)
}
