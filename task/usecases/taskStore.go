package usecases

import (
	"bergenTech/task"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type TaskStore struct {
	Tasks map[string]task.Task
}

func New() task.Usecase {
	return &TaskStore{
		Tasks: map[string]task.Task{},
	}
}

func (ts *TaskStore) WriteState(w io.Writer) {
	if err := json.NewEncoder(w).Encode(ts.Tasks); err != nil {
		log.Println(err)
	}
}

func (ts *TaskStore) AddTask(t task.Task) error {
	if t.Name == "" {
		return fmt.Errorf("name must be not empty")
	}
	ts.Tasks[t.Name] = t
	return nil
}

func (ts *TaskStore) AddSubtask(parent string, sb task.SubTask) error {
	if parent == "" {
		return fmt.Errorf("parent must be not empty")
	}
	if sb.Name == "" {
		return fmt.Errorf("subtaskName must be not empty")
	}

	if v, ok := ts.Tasks[parent]; ok {
		v.SubTasks[sb.Name] = sb
	}
	return nil
}

func (ts *TaskStore) AddCost(taskName, parentName string, c task.Cost) error {
	if taskName == "" || parentName == "" || c.Name == "" {
		return fmt.Errorf("cannot transfer empty names")
	}

	if t, ok := ts.Tasks[taskName]; ok {
		if st, ok := t.SubTasks[parentName]; ok {
			st.Costs[c.Name] = c
			ts.recountCostsForSubtask(taskName, parentName)
			ts.recountCostsForTask(taskName)
		}
	}
	return nil
}

func (ts *TaskStore) recountCostsForSubtask(taskName, parentName string) {
	var averageCost, totalCost, nums int

	for _, v := range ts.Tasks[taskName].SubTasks[parentName].Costs {
		nums++
		totalCost += v.Value
		averageCost = totalCost / nums
	}
	s := ts.Tasks[taskName].SubTasks[parentName]
	s.TotalCost = totalCost
	s.AverageCost = averageCost
	ts.Tasks[taskName].SubTasks[parentName] = s
}

func (ts *TaskStore) recountCostsForTask(taskName string) {
	var averageCost, totalCost, nums int

	for _, v := range ts.Tasks[taskName].SubTasks {
		nums++
		totalCost += v.TotalCost
		averageCost += v.AverageCost / nums
	}
	t := ts.Tasks[taskName]
	t.TotalCost = totalCost
	t.AverageCost = averageCost
	ts.Tasks[taskName] = t
}
