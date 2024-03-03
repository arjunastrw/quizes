package Handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"mini-project/Model"
)

// Definisikan jwtKey sebagai sebuah string rahasia
var jwtKey = []byte("secret_key")

// RegisterHandler adalah fungsi untuk menangani permintaan registrasi pengguna.
func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", 405)
			return
		}

		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", 500)
			return
		}

		// Ambil data pengguna dari permintaan
		nama := r.FormValue("nama")
		email := r.FormValue("email")
		password := r.FormValue("password")
		role := "user"

		// Tentukan peran default jika peran tidak diberikan
		if role == "" {
			role = "User"
		} else {
			// Validasi nilai peran yang diberikan (harus 'Admin' atau 'User')
			if role != "Admin" && role != "User" {
				http.Error(w, "Role isn't 'Admin' or 'User'", 400)
				return
			}
		}

		// Validasi input
		if nama == "" || email == "" || password == "" {
			http.Error(w, "Nama, email, dan password diperlukan", 400)
			return
		}

		// Hash password pengguna dengan bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Gagal menghash password", 500)
			return
		}

		// Buat instansiasi objek User baru
		newUser := Model.User{
			Nama:     nama,
			Email:    email,
			Password: string(hashedPassword),
			Role:     role,
		}

		// Simpan pengguna baru ke dalam database
		err = newUser.Save(db)
		if err != nil {
			http.Error(w, "Gagal mendaftarkan pengguna", 500)
			return
		}

		// Buat token JWT untuk pengguna yang baru terdaftar
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["user_id"] = newUser.Id
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token berlaku selama 24 jam

		// Tandatangani token JWT menggunakan secret key
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "Gagal membuat token JWT", 500)
			return
		}

		// Beri tanggapan kepada pengguna yang berhasil diregistrasi
		response := map[string]interface{}{
			"message": "Pengguna berhasil diregistrasi",
			"user_id": newUser.Id,
			"token":   tokenString,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Gagal membuat respons JSON", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}
