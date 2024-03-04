package UserHandlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"mini-project/Model"
	"net/http"
	"strconv"
)

func GetAllUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil informasi pengguna dari konteks
		user, exists := c.Get("Role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User information not found in context"})
			return
		}

		// Cek apakah pengguna adalah admin
		if user != "Admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins are allowed to access this resource"})
			return
		}

		// Dapatkan daftar user dari database
		users, err := Model.GetAllUser(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
func GetUserByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil informasi pengguna dari konteks
		userRole, exists := c.Get("Role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User information not found in context"})
			return
		}

		// Cek apakah pengguna adalah admin
		if userRole != "Admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins are allowed to access this resource"})
			return
		}

		// Ambil id pengguna dari parameter URL
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Dapatkan informasi pengguna dari database berdasarkan ID
		user, err := Model.GetUserById(db, userID)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
