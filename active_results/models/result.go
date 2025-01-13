package models

import (
	"github.com/google/uuid"
)

type Result struct {
	ExamID      uuid.UUID `json:"exam_id"`
	GroupID     string    `json:"group_id"`
	StudentID   string    `json:"student_id"`
	StudentName string    `json:"student_name"`
	Scores      []float64 `json:"scores"`
}
