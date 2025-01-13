package handlers

import (
	"context"
	"net/http"
	"result_publisher/active_results/repositories"

	"github.com/gin-gonic/gin"
)

type ExamHandler struct {
	Repo *repositories.ExamRepository
}

// GetExamByID retrieves an exam by its ID
func (h *ExamHandler) GetExamByID(c *gin.Context) {
	// Extract exam ID from URL parameters
	examID := c.Param("id")
	if examID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing exam ID in request"})
		return
	}

	// Fetch the exam from the repository
	exam, err := h.Repo.GetExamByID(context.Background(), examID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch exam: " + err.Error()})
		return
	}

	// Return the exam as JSON
	c.JSON(http.StatusOK, exam)
}
