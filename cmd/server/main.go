package main

import (
	"log"
	"net/http"

	"todo/handler"
)

func main() {
	mux := http.NewServeMux()
	handler := handler.NewTaskHandler()

	mux.HandleFunc("POST /task", handler.AddTask)
	mux.HandleFunc("GET /tasks", handler.GetTasks)
	mux.HandleFunc("GET /task/{id}", handler.GetTask)
	mux.HandleFunc("DELETE /task/{id}", handler.DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
