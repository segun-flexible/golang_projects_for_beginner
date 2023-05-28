package main

import (
	"fmt"
	"mytodo/helpers"
)

var appRunning bool = true

type TodoType struct {
	id        int64
	title     string
	completed bool
}

func addNewTodo(title string, todos *[]TodoType) {
	*todos = append(*todos, TodoType{int64(len(*todos) + 1), title, false})
}

func main() {

	todos := []TodoType{}

	//Keep Bringing The Menu While App Is Still Running
	var userInput int

	for appRunning {

		fmt.Println("Welcome To My Todo")
		fmt.Println("Choose From The Menu")
		fmt.Println("1. Create New Todo")
		fmt.Println("2. View All Todo(s)")
		fmt.Println("3. Quit App")
		fmt.Print("\n\n")

		//Read User Input
		fmt.Scanln(&userInput)

		if userInput == 1 {
			helpers.ClearScreen()

			var todoTitle string

			fmt.Println("Create New Todo By Inputing What You Want To Do!")
			helpers.ScanInput(&todoTitle)
			for todoTitle == "" {
				fmt.Println("Todo Title Cannot Be Blank!")
				helpers.ScanInput(&todoTitle)
			}

			addNewTodo(todoTitle, &todos)
			s := fmt.Sprintf("New Todo Created - Total (%v)", len(todos))
			fmt.Println(s)
			fmt.Print("\n\n")

		} else if userInput == 2 {
			helpers.ClearScreen()
			fmt.Println("************* Todo(s) *************")
			fmt.Println()

			for _, t := range todos {
				fmt.Println(t.id, ". ", t.title, t.completed)
			}

			fmt.Print("\n\n")
			fmt.Println("Enter Any Button To Go To Menu")
			var q string
			fmt.Scanln(&q)
			helpers.ClearScreen()

		} else {

			//Check If User Quit
			if userInput == 3 {
				break
			}

			helpers.ClearScreen()
			fmt.Println("*** Your Input Is Not Recognize")
			fmt.Print("\n\n")
		}

	}

	helpers.ClearScreen()
	fmt.Println("App Stopped")

}
