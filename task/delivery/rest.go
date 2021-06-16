package delivery

import (
	"bergenTech/task"
	"encoding/json"
	"log"
	"net/http"
)

type TaskStoreHandler struct {
	store task.Usecase
}

func New(usecase task.Usecase) *TaskStoreHandler {
	return &TaskStoreHandler{store: usecase}
}

func (tsh *TaskStoreHandler) TaskState(w http.ResponseWriter, r *http.Request) {
	tsh.store.WriteState(w)
}

func (tsh *TaskStoreHandler) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	t := task.Task{
		Name:     "",
		SubTasks: map[string]task.SubTask{},
	}
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := tsh.store.AddTask(t); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (tsh *TaskStoreHandler) AddSubtaskHandler(w http.ResponseWriter, r *http.Request) {
	nst := struct {
		ParentName  string `json:"parentName"`
		SubtaskName string `json:"subtaskName"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&nst); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := tsh.store.AddSubtask(
		nst.ParentName, task.SubTask{Name: nst.SubtaskName, Costs: map[string]task.Cost{}},
	); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (tsh *TaskStoreHandler) AddCostHandler(w http.ResponseWriter, r *http.Request) {
	ncost := struct {
		TaskName   string    `json:"taskName"`
		ParentName string    `json:"parentName"`
		Cost       task.Cost `json:"cost"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&ncost); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := tsh.store.AddCost(ncost.TaskName, ncost.ParentName, ncost.Cost); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
