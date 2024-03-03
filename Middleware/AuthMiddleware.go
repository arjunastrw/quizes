package Middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

var SecretKey = "inikuncinya"

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dapatkan token JWT dari header Authorization
		tokenString := c.GetHeader("Authorization")

		// Parse dan verifikasi token JWT
		claims, err := ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse or verify token"})
			c.Abort()
			return
		}

		// Ambil role pengguna dari klaim token
		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins are allowed to access this resource"})
			c.Abort()
			return
		}

		// Jika pengguna memiliki peran admin, lanjutkan pemrosesan berikutnya
		c.Next()
	}
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse dan verifikasi token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Lakukan validasi metode tanda tangan token di sini
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		// Kembalikan kunci rahasia yang digunakan untuk menandatangani token
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Periksa kesalahan saat parsing token
	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	// Ambil klaim dari token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}
