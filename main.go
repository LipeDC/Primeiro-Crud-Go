package main

import (
	"log"
	"net/http"

	"github.com/LipeDC/simple-go-mod/config"
	"github.com/LipeDC/simple-go-mod/handlers"
	"github.com/LipeDC/simple-go-mod/models"
	"github.com/gorilla/mux"
)

// É a função chamada quando a aplicação é iniciada
func main() {

	dbConnection := config.SetupDB()

	_, err := dbConnection.Exec(models.CreateTableQuery)

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	TaskHandler := handlers.NewTaskHandler(dbConnection)

	router.HandleFunc("/tasks", TaskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", TaskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", TaskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", TaskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
