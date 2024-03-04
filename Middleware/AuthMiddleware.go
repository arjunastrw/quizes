package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"mini-project/Model"
	"net/http"
	"strings"
)

var SecretKey = "inikuncinya"

func AuthMiddleware(ctx *gin.Context) {
	if ctx.FullPath() != "/api/v1/login" && ctx.FullPath() != "/api/v1/register" && ctx.FullPath() != "/api/v1/logout" && ctx.FullPath() != "/api/v1/get-quiz" {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse or verify token"})
			ctx.Abort()
			return
		}

		authArr := strings.Split(authHeader, " ")
		if len(authArr) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}

		if authArr[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid bearer token"})
		}

		tokenStr := authArr[1]

		var authClaim Model.AuthClaimJWT

		token, err := jwt.ParseWithClaims(tokenStr, &authClaim, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil {
			panic(err.Error())
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}

		ctx.Set("Role", authClaim.Role)
		log.Println(authClaim)
		log.Println(token)
		ctx.Next()
	}
}
