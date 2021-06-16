package main

import (
	"bergenTech/task/delivery"
	"bergenTech/task/usecases"
	"log"
	"net/http"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	tUcase := usecases.New()
	dts := delivery.New(tUcase)

	r.Get("/task", dts.TaskState)
	r.Post("/task/add", dts.AddTaskHandler)
	r.Post("/subtask/add", dts.AddSubtaskHandler)
	r.Post("/cost/add", dts.AddCostHandler)

	log.Println("server starts at :8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
