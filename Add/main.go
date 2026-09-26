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
	newTask := Task{Name: taskName, Completed: false}
	return append(tasks, newTask)
}

func main() {
	// Create scanner for reading input lines
	scanner := bufio.NewScanner(os.Stdin)
	
	// Read number of existing tasks
	scanner.Scan()
	existingTasksStr := scanner.Text()
	
	// Read task names
	scanner.Scan()
	taskNamesStr := scanner.Text()
	
	existingTasksCount, _ := strconv.Atoi(existingTasksStr)
	
	var tasks []Task
	
	// Create existing tasks
	for i := 1; i <= existingTasksCount; i++ {
		taskName := fmt.Sprintf("Existing Task %d", i)
		tasks = append(tasks, Task{Name: taskName, Completed: false})
	}
	
	// Add new tasks if any are provided
	if taskNamesStr != "" {
		taskNames := strings.Split(taskNamesStr, ",")
		for _, taskName := range taskNames {
			tasks = addTask(tasks, taskName)
		}
	}
	
	// Print all tasks
	for _, task := range tasks {
		fmt.Printf("Task: %s, Completed: %t\n", task.Name, task.Completed)
	}
	
	// Print total count
	fmt.Printf("Total tasks: %d\n", len(tasks))
}
