package storage

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func SaveTasks(fileName string, task []Task) error {
	data, _ := json.MarshalIndent(task, "", " ")
	return os.WriteFile(fileName, data, 0644)
}
