package Model

import (
	"database/sql"
	"log"
)

type User struct {
	Id       int
	Nama     string
	Password string
	Email    string
	Role     string
}

// Save adalah metode untuk menyimpan pengguna ke dalam database.
func (u *User) Save(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO user (nama, email, password, role) VALUES (?, ?, ?, ?)", u.Nama, u.Email, u.Password, u.Role)
	return err
}

// Get User By Email untuk login
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, nama, email, password, role FROM user WHERE email = ?", email).Scan(&user.Id, &user.Nama, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(db *sql.DB, id int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, nama, email, password, role FROM user WHERE id = ?", id).Scan(&user.Id, &user.Nama, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Println("Failed to get user information:", err)
		return nil, err
	}
	return &user, nil
}
