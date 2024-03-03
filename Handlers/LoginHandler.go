package Handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mini-project/Middleware" // Menambahkan impor package Middleware
	"mini-project/Model"
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
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["user_id"] = user.Id
		claims["name"] = user.Nama
		claims["email"] = user.Email
		claims["role"] = user.Role
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token berlaku selama 24 jam

		// Tandatangani token JWT menggunakan secret key
		tokenString, err := token.SignedString([]byte(Middleware.SecretKey))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate token"})
			return
		}

		// Beri tanggapan kepada pengguna yang berhasil login
		response := map[string]interface{}{
			"message": "Login successfuly",
			"token":   tokenString,
		}
		c.JSON(200, response)
	}
}
