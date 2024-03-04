package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini-project/Config"
	"mini-project/Handlers/QuestionHandlers"
	"mini-project/Handlers/QuizHandlers"
	"mini-project/Handlers/UserHandlers"
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
		UserHandlers.RegisterHandler(db)(c)
	})
	//route untuk login
	router.POST("/api/v1/login", func(c *gin.Context) {
		// Panggil handler login dan lewatkan konteks Gin
		UserHandlers.LoginHandler(db)(c)
	})
	//route untuk logout
	router.DELETE("/api/v1/logout", func(c *gin.Context) {
		// Panggil handler logout
		UserHandlers.LogoutHandler(c)
	})

	// 				route untuk User
	// 				route untuk get All User
	router.GET("/api/v1/user", UserHandlers.GetAllUserHandler(db))
	// route untuk get user by id
	router.GET("/api/v1/user/:id", UserHandlers.GetUserByIdHandler(db))

	//				route untuk quiz
	//route untuk membuat quiz
	router.POST("/api/v1/create-quiz", QuizHandlers.CreateQuizHandler(db))

	// route untuk get quiz
	router.GET("/api/v1/quiz", QuizHandlers.GetAllQuizHandler(db))

	// route untuk update quiz
	router.PUT("/api/v1/quiz/:id", QuizHandlers.UpdateQuizHandler(db))

	// Route untuk delete quiz
	router.DELETE("/api/v1/quiz/:id", QuizHandlers.DeleteQuizHandler(db))

	// Route untuk pertanyaan
	// Route untuk membuat pertanyaan dan jawabnnya
	router.POST("/api/v1/question", QuestionHandlers.CreateQuestionHandler(db))

	// route untuk get all pertanyaan dan jawbaan
	router.GET("/api/v1/question", QuestionHandlers.GetAllQuestionHandler(db))

	// Menjalankan server HTTP
	router.Run(":8080")
}
