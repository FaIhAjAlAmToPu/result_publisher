package repositories

import (
	"context"
	"errors"
	"result_publisher/active_results/models"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type ExamRepository struct {
	Session gocqlx.Session
}

// GetExamByID retrieves an exam by its ID
func (e *ExamRepository) GetExamByID(ctx context.Context, id string) (*models.Exam, error) {
	var exam models.Exam

	// Build and execute the query
	query := qb.Select("active_results.exam").
		Columns("id", "name", "separator_name", "held_by", "format", "publishing_date", "end_time").
		Where(qb.Eq("id")).
		Query(e.Session).BindMap(qb.M{"id": id})

	if err := query.WithContext(ctx).GetRelease(&exam); err != nil {
		return nil, errors.New("exam not found: " + err.Error())
	}

	return &exam, nil
}
