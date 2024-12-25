package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./courses.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create tables
	tables := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			role TEXT NOT NULL,
			bio TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS courses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			duration TEXT NOT NULL,
			price REAL NOT NULL,
			instructor TEXT NOT NULL,
			category TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS enrollments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			course_id INTEGER NOT NULL,
			completed BOOLEAN DEFAULT FALSE,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
			FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS lessons (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			course_id INTEGER NOT NULL,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			video_url TEXT,
			"order" INTEGER NOT NULL,
			FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS progress (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			enrollment_id INTEGER NOT NULL,
			lesson_id INTEGER NOT NULL,
			completed BOOLEAN DEFAULT FALSE,
			FOREIGN KEY (enrollment_id) REFERENCES enrollments (id) ON DELETE CASCADE,
			FOREIGN KEY (lesson_id) REFERENCES lessons (id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS reviews (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			course_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			rating INTEGER NOT NULL CHECK(rating >= 1 AND rating <= 5),
			comment TEXT,
			FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		)`,
	}

	for _, query := range tables {
		if _, err := DB.Exec(query); err != nil {
			log.Fatalf("Failed to execute query: %v", err)
		}
	}

	fmt.Println("Database connected and tables initialized.")
}
