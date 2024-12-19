package handlers

import (
	"myapp/models"
	"myapp/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResultHandler struct {
	Repo *repositories.ResultRepository
}

// GetResults fetches results based on optional user_id query parameter
func (h *ResultHandler) GetResults(c *gin.Context) {
	userIDParam := c.Query("user_id")

	if userIDParam != "" {
		userID, err := strconv.Atoi(userIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id parameter"})
			return
		}
		results, err := h.Repo.GetResultsByUserID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch results"})
			return
		}
		c.JSON(http.StatusOK, results)
		return
	}

	results, err := h.Repo.GetAllResults()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch results"})
		return
	}
	c.JSON(http.StatusOK, results)
}

// GetResultByID fetches a result by ID
func (h *ResultHandler) GetResultByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	result, err := h.Repo.GetResultByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Result not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// PublishResult creates a new result
func (h *ResultHandler) PublishResult(c *gin.Context) {
	var result models.Result
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.Repo.CreateResult(result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish result"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Result published successfully"})
}

// UpdateResult updates an existing result
func (h *ResultHandler) UpdateResult(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	var result models.Result
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	result.ID = id

	if err := h.Repo.UpdateResult(result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Result updated successfully"})
}

// DeleteResult deletes a result by ID
func (h *ResultHandler) DeleteResult(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	if err := h.Repo.DeleteResult(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete result"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Result deleted successfully"})
}
