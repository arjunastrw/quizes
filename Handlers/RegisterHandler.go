package Handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mini-project/Model"
	"net/http"
)

// RegisterHandler adalah fungsi untuk menangani permintaan registrasi pengguna.
func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser Model.User
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode JSON request body"})
			return
		}

		// validasi data
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

		err = newUser.SaveUser(db, &newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		response := gin.H{
			"message":   "Pengguna berhasil diregistrasi",
			"userEmail": newUser.Email,
		}
		c.JSON(http.StatusOK, response)
	}
}
