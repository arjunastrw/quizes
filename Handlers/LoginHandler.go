package Handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"mini-project/Middleware"
	"mini-project/Model"
	"net/http"
)

// LoginHandler adalah fungsi untuk menangani permintaan login pengguna.
func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			c.JSON(405, gin.H{"error": "Method not allowed"})
			return
		}

		// Parse form data
		err := c.Request.ParseForm()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to parse form data"})
			return
		}

		// Ambil data pengguna dari form
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")

		// Cari pengguna berdasarkan email
		user, err := Model.GetUserByEmail(db, email)
		if err != nil {
			c.JSON(500, gin.H{"error": "Please Contact Support"})
			return
		}

		// Jika pengguna tidak ditemukan, beri respons error
		if user == nil {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		// Bandingkan password yang diberikan dengan password yang disimpan di database
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		// Buat token JWT untuk pengguna yang berhasil login
		tokenClaim := Model.AuthClaimJWT{
			Role: user.Role,
		}

		// Tandatangani token JWT menggunakan secret key
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)

		tokenString, err := token.SignedString([]byte(Middleware.SecretKey))

		// Beri tanggapan kepada pengguna yang berhasil login
		response := map[string]interface{}{
			"message": "Login successfully",
			"token":   tokenString,
		}
		c.JSON(200, response)
	}
}
