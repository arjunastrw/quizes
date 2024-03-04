package Model

import "database/sql"

type Question struct {
	ID           int    `json:"id"`
	Pertanyaan   string `json:"pertanyaan"`
	OpsiJawaban  string `json:"opsi_jawaban"`
	JawabanBenar string `json:"jawaban_benar"`
	QuizID       int    `json:"id_quiz"`
}

func (q *Question) SaveQuestion(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO pertanyaan (pertanyaan, opsi_jawaban, jawaban_benar, id_quiz) VALUES (?, ?, ?, ?)", q.Pertanyaan, q.OpsiJawaban, q.JawabanBenar, q.QuizID)
	return err
}

func GetAllQuestion(db *sql.DB) ([]Question, error) {
	var questions []Question
	rows, err := db.Query("SELECT id, pertanyaan, opsi_jawaban, jawaban_benar, id_quiz FROM pertanyaan")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var question Question
		if err := rows.Scan(&question.ID, &question.Pertanyaan, &question.OpsiJawaban, &question.JawabanBenar, &question.QuizID); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}
