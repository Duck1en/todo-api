package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Duck1en/todo-api/storage"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := storage.GetAllTasks()
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	task := storage.GetTaskById(taskID)
	if task == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func CreateTask() {

}

func UpdateTask() {

}

func DeleteTask() {}
