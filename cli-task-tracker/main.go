package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

const filename = "tasks.json"

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
func loadTasks() ([]Task, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func addTask(title string) {
	tasks, _ := loadTasks()

	newTask := Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	err := saveTasks(tasks)
	if err != nil {
		return
	}
	fmt.Printf("Added New Task: %s\n", title)
}

func listTask() {
	tasks, _ := loadTasks()
	fmt.Printf("\n ----- Tasks List ----- \n")
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "X"
		}

		timeStr := task.CreatedAt.Format("Jan 02 03:04 PM")
		fmt.Printf("[%s] %d: %s [Created At: %s]\n", status, task.ID, task.Title, timeStr)
	}
}

func promptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	// ReadString reads until it hits the 'Enter' Key
	input, _ := reader.ReadString('\n')

	//Clean Up the input
	return strings.TrimSpace(input)
}

func completeTask(idStr string) {
	// 1. Convert String to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID ... Plese enter a valid integer")
		return
	}

	// 2. Load Task
	tasks, _ := loadTasks()
	found := false

	// 3. Loop through to find appropriate task
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if found {
		err := saveTasks(tasks)
		if err != nil {
			return
		}
		fmt.Printf("\n Completed Task with id %s\n", idStr)
	} else {
		fmt.Printf("\n Could not find task with ID %d\n", id)
	}
}

func deleteTask(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID ... Plese enter a valid integer")
		return
	}

	tasks, _ := loadTasks()
	var newTask []Task
	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		newTask = append(newTask, t)
	}

	if found {
		err := saveTasks(newTask)
		if err != nil {
			return
		}
		fmt.Printf("\n Deleted Task with id %d\n", id)
	} else {
		fmt.Printf("\n Could not find task with ID %d\n", id)
	}
}

func main() {
	if len(os.Args) < 2 {
		for {
			fmt.Println("\n1. ADD TASK | 2. LIST TASKS | 3. MARK AS COMPLETE | 4. Delete Task | 5. EXIT")
			choice := promptInput("Chose an option: ")

			switch choice {
			case "1":
				title := promptInput("Enter Task Title: ")
				if title != "" {
					addTask(title)
				}
			case "2":
				listTask()
			case "3":
				id := promptInput("Enter Task ID: ")
				completeTask(id)
			case "4":
				deleteTask(promptInput("Enter Task ID: "))
			case "5", "EXIT", "quit", "Exit", "exit":
				fmt.Println("Bye!")
				return
			default:
				fmt.Println("Unknown Option!")
			}
		}
	}
}
