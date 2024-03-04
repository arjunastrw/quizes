package QuestionHandlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"mini-project/Model"
	"net/http"
)

func CreateQuestionHandler(db *sql.DB) gin.HandlerFunc {
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

		var newQuestion Model.Question
		if err := c.BindJSON(&newQuestion); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body"})
			return
		}

		// Validasi data validasi pertanyaan
		if newQuestion.Pertanyaan == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Pertanyaan diperlukan"})
			return
		}
		// validasi data opsi jawaban
		if newQuestion.OpsiJawaban == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Opsi Jawaban diperlukan"})
			return
		}

		// validasi data opsi jawaban
		if newQuestion.JawabanBenar == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Opsi Jawaban diperlukan"})
			return
		}

		if newQuestion.QuizID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Quiz ID diperlukan"})
			return
		}

		// Simpan pertanyaan baru ke dalam database
		err := newQuestion.SaveQuestion(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan pertanyaan baru"})
			return
		}
	}
}

func GetAllQuestionHandler(db *sql.DB) gin.HandlerFunc {
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

		questions, err := Model.GetAllQuestion(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil semua pertanyaan"})
			return
		}
		c.JSON(http.StatusOK, questions)
	}
}
