package repositories

import (
	"database/sql"
	"myapp/models"
)

type ResultRepository struct {
	DB *sql.DB
}

// GetAllResults fetches all results from the database
func (r *ResultRepository) GetAllResults() ([]models.Result, error) {
	query := "SELECT id, user_id, exam_id, score, grade, created_at FROM results"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Result
	for rows.Next() {
		var result models.Result
		if err := rows.Scan(&result.ID, &result.UserID, &result.ExamID, &result.Score, &result.Grade, &result.CreatedAt); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

// GetResultsByUserID fetches results by a specific user ID
func (r *ResultRepository) GetResultsByUserID(userID int) ([]models.Result, error) {
	query := "SELECT id, user_id, exam_id, score, grade, created_at FROM results WHERE user_id = $1"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Result
	for rows.Next() {
		var result models.Result
		if err := rows.Scan(&result.ID, &result.UserID, &result.ExamID, &result.Score, &result.Grade, &result.CreatedAt); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

// GetResultByID fetches a result by its ID
func (r *ResultRepository) GetResultByID(id int) (*models.Result, error) {
	query := "SELECT id, user_id, exam_id, score, grade, created_at FROM results WHERE id = $1"
	row := r.DB.QueryRow(query, id)

	var result models.Result
	if err := row.Scan(&result.ID, &result.UserID, &result.ExamID, &result.Score, &result.Grade, &result.CreatedAt); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateResult inserts a new result into the database
func (r *ResultRepository) CreateResult(result models.Result) error {
	query := "INSERT INTO results (user_id, exam_id, score, grade, created_at) VALUES ($1, $2, $3, $4, NOW())"
	_, err := r.DB.Exec(query, result.UserID, result.ExamID, result.Score, result.Grade)
	return err
}

// UpdateResult updates an existing result by ID
func (r *ResultRepository) UpdateResult(result models.Result) error {
	query := "UPDATE results SET user_id = $1, exam_id = $2, score = $3, grade = $4 WHERE id = $5"
	_, err := r.DB.Exec(query, result.UserID, result.ExamID, result.Score, result.Grade, result.ID)
	return err
}

// DeleteResult deletes a result by ID
func (r *ResultRepository) DeleteResult(id int) error {
	query := "DELETE FROM results WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
