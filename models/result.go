package models

type Result struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	ExamID    int     `json:"exam_id"`
	Score     float64 `json:"score"`
	Grade     string  `json:"grade"`
	CreatedAt string  `json:"created_at"`
}
