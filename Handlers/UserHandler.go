package Handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"mini-project/Model"
	"net/http"
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

func GetUserHandlerByNama(db *sql.DB) gin.HandlerFunc {
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

		// Dapatkan nama peserta dari database
		nama := c.Param("nama")
		user, err := Model.GetUserByNama(db, nama)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
