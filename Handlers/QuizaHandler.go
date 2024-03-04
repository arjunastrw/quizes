package Handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"mini-project/Model"
)

// CreateQuizHandler adalah fungsi untuk menangani permintaan pembuatan quiz baru oleh admin.
func CreateQuizHandler(db *sql.DB) gin.HandlerFunc {
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

		// Parse request body
		var newQuiz Model.Quiz
		if err := c.BindJSON(&newQuiz); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body"})
			return
		}

		// Validasi data
		if newQuiz.Judul == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Judul quiz diperlukan"})
			return
		}

		// Set waktu mulai quiz
		newQuiz.WaktuMulai = time.Now()

		// Simpan quiz baru ke dalam database
		err := newQuiz.Save(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan quiz baru"})
			return
		}

		// Beri respons sukses
		c.JSON(http.StatusCreated, newQuiz)
	}
}

func GetAllQuizHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dapatkan daftar kuis dari database
		quizes, err := Model.GetAllQuiz(db)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": "Failed to get quizes"})
			return
		}

		// Beri respons dengan daftar kuis
		c.JSON(http.StatusOK, quizes)
	}
}
