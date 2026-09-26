package main
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"

)

type Task struct {
	Name      string
	Completed bool
}

func viewAllTasks(t []Task) {
	var task string

	for _, oneTask := range t{
		if oneTask.Completed{
			task = fmt.Sprintf("[x] %s", oneTask.Name)
			fmt.Println(task)
		}else{
			task = fmt.Sprintf("[ ] %s", oneTask.Name)
			fmt.Println(task)
		}
	}
	
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskNum := scanner.Text()

	scanner.Scan()
	taskNameComb := scanner.Text()

	taskInt, _ := strconv.Atoi(taskNum)
	
	tasks := strings.Split(taskNameComb, ",")

	allTasks := []Task{}

	for _, task := range tasks{
		taskDesc := strings.Split(task, ":")
		isCompleted,_ := strconv.ParseBool(taskDesc[1]) 
		newTask := Task{Name: taskDesc[0], Completed: isCompleted}
		allTasks = append(allTasks, newTask)

	}
	viewAllTasks(allTasks)

	// summary
	completed_count := 0
	for _, task := range allTasks{
		if task.Completed{
			completed_count++
		}
	}
	incomplete_count:= taskInt- completed_count
	fmt.Printf("Total: %d tasks (%d completed, %d remaining)", taskInt, completed_count,incomplete_count)

}