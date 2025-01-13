package handlers

import (
	"context"
	"net/http"
	"result_publisher/active_results/models"

	"result_publisher/active_results/repositories"

	"github.com/gin-gonic/gin"
)

type ResultHandler struct {
	Repo *repositories.ResultRepository
}

// GetResultByStudentID retrieves a result by exam ID, group ID, and student ID
func (h *ResultHandler) GetResultByStudentID(c *gin.Context) {
	// Extract parameters from the URL
	examID := c.Param("exam_id")
	groupID := c.Param("group_id")
	studentID := c.Param("student_id")

	// Validate input if needed (e.g., UUID check for examID)
	// Example: validate examID as UUID

	// Fetch the result from the repository
	result, err := h.Repo.GetResultByStudentID(context.Background(), examID, groupID, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch result: " + err.Error()})
		return
	}

	// Return the result as JSON
	c.JSON(http.StatusOK, result)
}

// InsertResult handles inserting a new result into the database
func (h *ResultHandler) InsertResult(c *gin.Context) {
	var result models.Result

	// Decode JSON payload
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Insert result into the repository
	err := h.Repo.InsertResult(context.Background(), result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert result: " + err.Error()})
		return
	}

	// Respond with success
	c.JSON(http.StatusCreated, gin.H{"message": "Result inserted successfully"})
}
