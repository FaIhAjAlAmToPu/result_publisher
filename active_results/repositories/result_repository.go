package repositories

import (
	"context"
	"result_publisher/active_results/models"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type ResultRepository struct {
	Session gocqlx.Session
}

// GetResultByStudentID fetches a result by exam ID, group ID, and student ID
func (r *ResultRepository) GetResultByStudentID(ctx context.Context, examID, groupID, studentID string) (models.Result, error) {
	var result models.Result

	// Use qb to build the SELECT query
	query := qb.Select("active_results.result").
		Columns("exam_id", "group_id", "student_id", "student_name", "scores").
		Where(qb.Eq("exam_id"), qb.Eq("group_id"), qb.Eq("student_id")).
		Query(r.Session)

	// Bind the parameters and execute the query
	err := query.BindMap(qb.M{
		"exam_id":    examID,
		"group_id":   groupID,
		"student_id": studentID,
	}).GetRelease(&result)

	return result, err
}

// InsertResult inserts a new result into the database
func (r *ResultRepository) InsertResult(ctx context.Context, result models.Result) error {
	// Use qb to build the INSERT query
	query := qb.Insert("active_results.result").
		Columns("exam_id", "group_id", "student_id", "student_name", "scores").
		Query(r.Session)

	// Bind the parameters and execute the query
	return query.BindStruct(result).Exec()
}
