package auth

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Auth token not informed",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		token, err := ValidadeToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("UserID", claims["sub"])

		ctx.Next()

	}
}
