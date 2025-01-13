package models

import (
	"time"

	"github.com/google/uuid"
)

type Exam struct {
	ID             uuid.UUID          `json:"id"`
	Name           string             `json:"name"`
	SeparatorName  string             `json:"separator_name"`
	HeldBy         string             `json:"held_by"`
	Format         map[string]float64 `json:"format"`
	PublishingDate time.Time          `json:"publishing_date"`
	EndTime        time.Time          `json:"end_time"`
}
