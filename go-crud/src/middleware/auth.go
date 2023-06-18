package middleware

import (
	"go-crud/src/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		parts := strings.Split(context.GetHeader("Authorization"), "Bearer ")
		if len(parts) == 2 {
			token := parts[1]
			user, err := services.ValidateToken(token)
			if err != nil {
				context.JSON(401, gin.H{"status": 401, "message": "Unauthorized"})
				context.Abort()
				return
			}
			context.Request.Header.Set("email", user.Email)
			context.Next()

		} else {
			context.JSON(401, gin.H{"status": 401, "message": "Unauthorized"})
			context.Abort()
			return
		}

	}
}
