package Model

import (
	"database/sql"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// Save user Register
// Save adalah fungsi untuk menyimpan data pengguna ke database.
func (s *User) SaveUser(db *sql.DB, user *User) error {
	_, err := db.Exec("INSERT INTO user (nama, email, password, role) VALUES (?, ?, ?, ?)",
		user.Nama, user.Email, user.Password, user.Role)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAllUser(db *sql.DB) ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT id, nama, email, role FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Nama, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Get User By Email untuk login
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, nama, email, password, role FROM user WHERE email = ?", email).Scan(&user.Id, &user.Nama, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	log.Println(user)
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
