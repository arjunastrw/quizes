package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini-project/Config"
	"mini-project/Handlers"
)

func main() {
	//koneksi ke database
	db := Config.LoadConfig()
	defer db.Close()
	fmt.Println("Connected to Database!")

	// Membuat instance Gin Engine
	router := gin.Default()

	// Menambahkan rute untuk endpoint registrasi
	//route untuk register
	router.POST("/api/v1/register", func(c *gin.Context) {
		// Panggil handler RegisterHandler dan lewatkan konteks Gin
		Handlers.RegisterHandler(db)
	})
	//route untuk login
	router.POST("/api/v1/login", func(c *gin.Context) {
		// Panggil handler login dan lewatkan konteks Gin
		Handlers.LoginHandler(db)(c)
	})
	//route untuk logout
	router.DELETE("/api/v1/logout", func(c *gin.Context) {
		// Panggil handler logout
		Handlers.LogoutHandler(c)
	})

	//route untuk membuat quiz
	router.POST("/create-quiz", Handlers.CreateQuizHandler(db))

	// route untuk get quiz
	router.GET("/get-quiz", Handlers.GetAllQuiz(db))

	// Menjalankan server HTTP
	router.Run(":8080")
}
