package Config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func LoadConfig() *sql.DB {
	// Buat koneksi ke database
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/quiz_app")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Uji koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	return db
}
