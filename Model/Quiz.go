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
		var waktuMulaiStr, waktuSelesaiStr string
		if err := rows.Scan(&quiz.ID, &quiz.Judul, &quiz.Deskripsi, &waktuMulaiStr, &waktuSelesaiStr); err != nil {
			return nil, err
		}
		// Mengonversi string waktu menjadi tipe data time.Time
		layout := "2006-01-02 15:04:05" // Format yang sesuai dengan format di database
		quiz.WaktuMulai, err = time.Parse(layout, waktuMulaiStr)
		if err != nil {
			return nil, err
		}
		quiz.WaktuSelesai, err = time.Parse(layout, waktuSelesaiStr)
		if err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}
	return quizzes, nil
}

func UpdateQuiz(db *sql.DB, id int, judul, deskripsi string) error {
	_, err := db.Exec("UPDATE quiz SET judul = ?, deskripsi = ? WHERE id = ?", judul, deskripsi, id)
	return err
}
