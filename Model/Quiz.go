package Model

import (
	"database/sql"
	"time"
)

// Quiz adalah struktur data untuk tabel quiz dalam database.
type Quiz struct {
	ID           int       `json:"id"`
	Judul        string    `json:"judul"`
	Deskripsi    string    `json:"deskripsi"`
	WaktuMulai   time.Time `json:"waktu_mulai"`
	WaktuSelesai time.Time `json:"waktu_selesai"`
}

// Save menyimpan data quiz ke dalam database.
func (q *Quiz) Save(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO quiz (judul, deskripsi, waktu_mulai, waktu_selesai) VALUES (?, ?, ?, ?)", q.Judul, q.Deskripsi, q.WaktuMulai, q.WaktuSelesai)
	return err
}

// GetAllQuiz mengambil semua data quiz dari database.
func GetAllQuiz(db *sql.DB) ([]Quiz, error) {
	var quizzes []Quiz
	rows, err := db.Query("SELECT id, judul, deskripsi, waktu_mulai, waktu_selesai FROM quiz")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var quiz Quiz
		if err := rows.Scan(&quiz.ID, &quiz.Judul, &quiz.Deskripsi, &quiz.WaktuMulai, &quiz.WaktuSelesai); err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}
	return quizzes, nil
}
