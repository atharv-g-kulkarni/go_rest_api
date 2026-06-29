package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/atharv-g-kulkarni/go_rest_api/utils"
)

func Authenticate(context *gin.Context){
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}