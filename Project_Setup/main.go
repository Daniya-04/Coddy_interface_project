package main

import "fmt"

// Define the Task struct
type Task struct{
	Name string
	Completed bool
}

func main() {
    // Read input
    var appName string
    var userName string
    fmt.Scanln(&appName)
    fmt.Scanln(&userName)
    
    // Create a slice of Task structs
	Tasks := []Task{}

    // Display welcome message and menu
	fmt.Printf("Welcome to %s,%s!\n", appName, userName)
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Remove Task")
	fmt.Println("5. Exit")
	// Print current tasks count
    fmt.Printf("Current tasks: %d", len(Tasks))
}