package Handlers

import (
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mini-project/Middleware"
	"mini-project/Model"
)

// RegisterHandler adalah fungsi untuk menangani permintaan registrasi pengguna.
func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser Model.User
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode JSON request body"})
			return
		}

		if newUser.Nama == "" || newUser.Email == "" || newUser.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nama, email, dan password diperlukan"})
			return
		}

		if newUser.Role == "" {
			newUser.Role = "User"
		} else if newUser.Role != "Admin" && newUser.Role != "User" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Role isn't 'Admin' or 'User'"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		newUser.Password = string(hashedPassword)

		err = newUser.Save(db * sql.DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": newUser.Id,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString(Middleware.SecretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token"})
			return
		}

		response := gin.H{
			"message": "Pengguna berhasil diregistrasi",
			"user_id": newUser.Id,
			"token":   tokenString,
		}
		c.JSON(http.StatusOK, response)
	}
}
