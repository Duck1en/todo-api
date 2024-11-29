package storage

import (
	"database/sql"
	"log"

	"github.com/Duck1en/todo-api/models"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDB(databasePath string) {
	var err error
	db, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatalf("Failed to connnect to the database: %v", err)
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        completed BOOLEAN NOT NULL
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create tasks in table: %v", err)
	}
}

func GetAllTasks() []models.Task {
	rows, err := db.Query("SELECT id, title, completed FROM tasks")
	if err != nil {
		log.Printf("Failed to fetch tasks: %v", err)
		return nil
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Competed)
		if err != nil {
			log.Printf("Failed to scan task row: %v", err)
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func GetTaskById(id int) *models.Task {
	query := "SELECT id, title, completed FROM tasks WHERE id = ?"
	row := db.QueryRow(query, id)

	var task models.Task
	err := row.Scan(&task.ID, &task.Title, &task.Competed)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		log.Printf("Failed to fetch task by ID: %v", err)
		return nil
	}
	return &task
}

func AddTask(title string) models.Task {
	query := "INSERT INTO tasks (title, completed) VALUES (?, ?)"
	result, err := db.Exec(query, title, false)
	if err != nil {
		log.Printf("Failed to add task: %v", err)
		return models.Task{}
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error retrieving ID for newly created task: %v", err)
	}
	return models.Task{
		ID:       int(id),
		Title:    title,
		Competed: false,
	}
}

func UpdateTask(id int, completed bool) *models.Task {
	query := "UPDATE tasks SET completed = ? WHERE id = ?"
	_, err := db.Exec(query, completed, id)
	if err != nil {
		log.Printf("Failed to update task: %v", err)
		return nil
	}

	return GetTaskById(id)
}

func DeleteTask(id int) bool {
	query := "DELETE FROM tasks WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete task: %v", err)
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to retrieve rows affected: %v", err)
		return false
	}
	return rowsAffected > 0
}
