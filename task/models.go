package task

type Task struct {
	Name string `json:"name"`

	SubTasks map[string]SubTask `json:"subtasks,omitempty"`

	TotalCost   int `json:"totalCost"`
	AverageCost int `json:"averageCost"`
}

type SubTask struct {
	Name  string          `json:"name"`
	Costs map[string]Cost `json:"costs,omitempty"`

	TotalCost   int `json:"totalCost"`
	AverageCost int `json:"averageCost"`
}

type Cost struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
