package main

import (
	"fmt"
	"result_publisher/active_results/database"
	"result_publisher/active_results/handlers"
	"result_publisher/active_results/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	session := database.ConnectToCluster()
	defer func() {
		session.Close()
		fmt.Println("Database connection closed.")
	}()

	// database.CreateKeyspace(session)
	database.CreateTables(session)

	// Initialize repository and handlers
	resultRepo := &repositories.ResultRepository{Session: session}
	resultHandler := &handlers.ResultHandler{Repo: resultRepo}
	examRepo := &repositories.ExamRepository{Session: session}
	examHandler := &handlers.ExamHandler{Repo: examRepo}

	// // Initialize Gin router
	r := gin.Default()

	// // Route definitions
	r.GET("/exams/:id", examHandler.GetExamByID)
	r.GET("/exams/:exam_id/results/:group_id/:student_id", resultHandler.GetResultByStudentID)
	// r.GET("/results/:id", resultHandler.GetResultByID)
	// r.POST("/results", resultHandler.PublishResult)
	// r.PUT("/results/:id", resultHandler.UpdateResult)
	// r.DELETE("/results/:id", resultHandler.DeleteResult)

	// // Route to run migrations
	// r.GET("/migrate", func(c *gin.Context) {
	// 	fmt.Println("Running migrations...")
	// 	if db == nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
	// 		return
	// 	}
	// 	populate.RunMigrations(db)
	// 	c.JSON(http.StatusOK, gin.H{"message": "Database tables created successfully!"})
	// })

	// // Route to insert sample data
	// r.GET("/populate", func(c *gin.Context) {
	// 	fmt.Println("Inserting sample data...")
	// 	if db == nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
	// 		return
	// 	}
	// 	populate.SampleInserts(db)
	// 	c.JSON(http.StatusOK, gin.H{"message": "Sample data inserted successfully!"})
	// })

	// // Test routes
	// r.GET("/", func(c *gin.Context) {
	// 	fmt.Println("Root endpoint hit")
	// 	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Result Publisher!"})
	// })

	// // Handlers for result operations
	// r.GET("/results", func(c *gin.Context) {
	// 	handlers.GetResults(c, db)
	// })

	// r.POST("/results", func(c *gin.Context) {
	// 	handlers.PublishResult(c, db)
	// })

	// // Graceful shutdown handling
	// go func() {
	// 	r.Run(":8080")
	// }() //started in a separate goroutine

	// quit := make(chan os.Signal, 1)                    //a channel to recv os signal
	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM) //reg this channe; to listen for specific signals
	// <-quit                                             //block main thread until signal recvd
	// fmt.Println("Shutting down server...")
}
