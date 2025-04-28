package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

const todoFile = "todos.json"

func main() {
	todos := loadTodos()

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  add <task description>  - Add a new task")
		fmt.Println("  list                    - List all tasks")
		fmt.Println("  done <task number>      - Mark a task as done")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		description := strings.Join(os.Args[2:], " ")
		todos = addTodo(todos, description)
		saveTodos(todos)
		fmt.Println("Task added.")

	case "list":
		listTodos(todos)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task number to mark as done.")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil || index < 1 || index > len(todos) {
			fmt.Println("Invalid task number.")
			return
		}
		todos = markDone(todos, index-1)
		saveTodos(todos)
		fmt.Println("Task marked as done.")

	default:
		fmt.Println("Unknown command:", command)
	}
}

func loadTodos() []Todo {
	file, err := os.ReadFile(todoFile)
	if err != nil {
		return []Todo{}
	}
	var todos []Todo
	err = json.Unmarshal(file, &todos)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return []Todo{}
	}
	return todos
}

func saveTodos(todos []Todo) {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	err = os.WriteFile(todoFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func addTodo(todos []Todo, description string) []Todo {
	todo := Todo{
		Description: description,
		Done:        false,
	}
	return append(todos, todo)
}

func listTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for i, todo := range todos {
		status := " "
		if todo.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, todo.Description)
	}
}

func markDone(todos []Todo, index int) []Todo {
	todos[index].Done = true
	return todos
}
