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

// Save user Register
// Save adalah fungsi untuk menyimpan data pengguna ke database.
func Save(db *sql.DB, user *User) error {
	_, err := db.Query("INSERT INTO users (nama, email, password, role) VALUES (?, ?, ?, ?)",
		user.Nama, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
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

func GetUserByNama(db *sql.DB, nama string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, nama, email, password, role FROM user WHERE nama = ?", nama).Scan(&user.Id, &user.Nama, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Println("Failed to get user information:", err)
		return nil, err
	}
	return &user, nil
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
