package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func addTask(tasks []Task, taskName string) []Task {
	return append(tasks, Task{Name: taskName, Completed: false})
}

func viewAllTasks(tasks []Task) {
	for _, task := range tasks {
		if task.Completed {
			fmt.Printf("[x] %s\n", task.Name)
		} else {
			fmt.Printf("[ ] %s\n", task.Name)
		}
	}
}

func completeTask(tasks *[]Task, index int) {
	(*tasks)[index].Completed = true
}

func removeTask(tasks []Task, index int) []Task {
	return append(tasks[:index], tasks[index+1:]...)
}

func printSummary(tasks []Task) {
	completed := 0
	for _, task := range tasks {
		if task.Completed {
			completed++
		}
	}
	fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", len(tasks), completed, len(tasks)-completed)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	names := strings.Split(scanner.Text(), ",")
	appName := names[0]
	userName := names[1]

	scanner.Scan()
	taskData := scanner.Text()

	scanner.Scan()
	actionData := scanner.Text()

	var tasks []Task
	if taskData != "" {
		for _, entry := range strings.Split(taskData, ",") {
			parts := strings.Split(entry, ":")
			tasks = append(tasks, Task{Name: parts[0], Completed: parts[1] == "true"})
		}
	}

	fmt.Printf("Welcome to %s, %s!\n", appName, userName)
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Remove Task")
	fmt.Println("5. Exit")
	fmt.Printf("Current tasks: %d\n", len(tasks))

	for _, action := range strings.Split(actionData, ",") {
		parts := strings.Split(action, "|")
		name := parts[0]
		param := ""
		if len(parts) > 1 {
			param = parts[1]
		}

		switch name {
		case "add":
			fmt.Println("--- ADD TASK ---")
			tasks = addTask(tasks, param)
			fmt.Printf("Task '%s' added!\n", param)

		case "view":
			fmt.Println("--- VIEW TASKS ---")
			viewAllTasks(tasks)
			printSummary(tasks)

		case "complete":
			fmt.Println("--- COMPLETE TASK ---")
			index, _ := strconv.Atoi(param)
			if index < 0 || index >= len(tasks) {
				fmt.Println("Invalid task number")
			} else {
				completeTask(&tasks, index)
				fmt.Printf("Task '%s' marked as completed!\n", tasks[index].Name)
			}

		case "remove":
			fmt.Println("--- REMOVE TASK ---")
			index, _ := strconv.Atoi(param)
			if index < 0 || index >= len(tasks) {
				fmt.Println("Invalid task number")
			} else {
				removedName := tasks[index].Name
				tasks = removeTask(tasks, index)
				fmt.Printf("Task '%s' removed successfully!\n", removedName)
			}

		case "exit":
			fmt.Println("--- EXIT ---")
			fmt.Println("Final list:")
			viewAllTasks(tasks)
			printSummary(tasks)
			fmt.Printf("Goodbye, %s!\n", userName)
			return
		}
	}
}
