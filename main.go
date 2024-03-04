package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini-project/Config"
	"mini-project/Handlers"
	"mini-project/Middleware"
)

func main() {
	//koneksi ke database
	db := Config.LoadConfig()
	defer db.Close()
	fmt.Println("Connected to Database!")

	// Membuat instance Gin Engine
	router := gin.Default()
	router.Use(Middleware.AuthMiddleware)

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

	// 				route untuk User
	// 				route untuk get All User
	router.GET("/api/v1/user", Handlers.GetAllUserHandler(db))

	//				route untuk get User By Nama
	router.GET("/api/v1/user/:nama", Handlers.GetUserHandlerByNama(db))

	//				route untuk quiz
	//route untuk membuat quiz
	router.POST("/api/v1/create-quiz", Handlers.CreateQuizHandler(db))

	// route untuk get quiz
	router.GET("/api/v1/quiz", Handlers.GetAllQuizHandler(db))

	// Menjalankan server HTTP
	router.Run(":8080")
}
