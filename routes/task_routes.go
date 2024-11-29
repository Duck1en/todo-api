package routes

import (
	"github.com/Duck1en/todo-api/controllers"
	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id:[0-9]+}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id:[0-9]+}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id:[0-9]+}", controllers.DeleteTask).Methods("DELETE")
}
