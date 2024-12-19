package populate

import (
	"database/sql"
	"fmt"
	"log"
)

// RunMigrations creates the necessary tables
func RunMigrations(db *sql.DB) {
	// Define the SQL statements for table creation
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(150) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(20) DEFAULT 'student',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	createExamsTable := `
	CREATE TABLE IF NOT EXISTS exams (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		subject VARCHAR(100) NOT NULL,
		exam_date DATE NOT NULL,
		max_marks INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	createResultsTable := `
	CREATE TABLE IF NOT EXISTS results (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		exam_id INT NOT NULL REFERENCES exams(id) ON DELETE CASCADE,
		score NUMERIC(5, 2) NOT NULL,
		grade VARCHAR(5) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(user_id, exam_id)
	);`

	// Execute each table creation query
	tables := []string{createUsersTable, createExamsTable, createResultsTable}

	for _, tableQuery := range tables {
		_, err := db.Exec(tableQuery)
		if err != nil {
			log.Fatalf("Error creating table: %v", err)
		}
		fmt.Println("Table created successfully!")
	}
}

// SampleInserts populates the database with sample data
func SampleInserts(db *sql.DB) {
	// Sample data insertion
	usersInsert := `
	INSERT INTO users (name, email, password, role) VALUES
	('Alice', 'alice@example.com', 'password123', 'student'),
	('Bob', 'bob@example.com', 'password123', 'student'),
	('Admin', 'admin@example.com', 'admin123', 'admin')
	ON CONFLICT DO NOTHING;`

	examsInsert := `
	INSERT INTO exams (name, subject, exam_date, max_marks) VALUES
	('Math Midterm', 'Mathematics', '2024-12-01', 100),
	('Science Final', 'Science', '2024-12-15', 100)
	ON CONFLICT DO NOTHING;`

	resultsInsert := `
	INSERT INTO results (user_id, exam_id, score, grade) VALUES
	(1, 1, 95.5, 'A'),
	(2, 2, 85.0, 'B')
	ON CONFLICT DO NOTHING;`

	// Execute the sample inserts
	sampleQueries := []string{usersInsert, examsInsert, resultsInsert}

	for _, query := range sampleQueries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Error inserting sample data: %v", err)
		}
		fmt.Println("Sample data inserted successfully!")
	}
}
