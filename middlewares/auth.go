package middlewares

import (
	"net/http"

	"example.com/plan-my-event/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// extract token from header
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to access the resource"})
		return
	}

	// check if the token is valid or not
	userID, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to access the resource"})
		return
	}
	context.Set("userID", userID)
	context.Next()
}
