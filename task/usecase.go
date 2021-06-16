package task

import "io"

type Usecase interface {
	AddTask(t Task) error
	AddSubtask(parentName string, st SubTask) error
	AddCost(taskName, parentName string, c Cost) error

	WriteState(w io.Writer)
}
