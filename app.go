package main

import (
	"fmt"
	"guess_game/helpers"
	"math/rand"
	"time"
)

var number int
var tries int = 10
var numberOfTries int = 0
var numberOfWins int64 = 0

func generateNumber() int {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)
	return random.Intn(100) + 1
}

func main() {

	var selectRunning bool = true

	for selectRunning {

		fmt.Print("\n\n")
		fmt.Println("****** Welcome To Guess Number Game ******")

		var startGame string
		if numberOfWins >= 1 {
			startGame = "1. Play Again"
		} else {
			startGame = "1. Play Now"
		}

		fmt.Println(startGame)
		fmt.Println("2. Read Rules")
		fmt.Println("3. Exit")
		fmt.Print("Enter Your Choice : ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {

			helpers.ClearScreen()
			number = generateNumber()
			fmt.Println("Number Generated: ***")
			fmt.Println("Now Guess The Number By Entering It")

			var userInput int
			for numberOfTries <= 9 {
				fmt.Println("Enter Your Number: ", "Your Tries Remain ", tries-numberOfTries)
				fmt.Scanln(&userInput)

				//Check if input is number
				if userInput <= 0 {
					userInput = 0
					continue
				}

				if userInput == number {
					numberOfWins++
					helpers.ClearScreen()
					fmt.Print("\n\n")
					fmt.Println("\U0001F3C6 Congratulations You Won \U0001F3C5")
					break
				} else if userInput > number {
					fmt.Print("\n\n")
					fmt.Println("Your Number Is Too High")
					numberOfTries++
				} else {
					fmt.Print("\n\n")
					fmt.Println("Your Number Is Too Low")
					numberOfTries++
				}

				//Each Input Is Turn Back To 0
				//So That I Can Check If User Enter Alphabet
				userInput = 0
			}

			//Failed
			numberOfTries = 0

			if userInput != number {
				helpers.ClearScreen()
				fmt.Println("****** Game Over, Your Tries Has Completed ******")
			}
			fmt.Print("\n\n")

		} else if choice == 2 {

			helpers.ClearScreen()

			fmt.Println("****** Rules ******")
			fmt.Println("1. Number is between 1 to 100")
			fmt.Println("2. Number of tries is 10")
			fmt.Println("3. Number of tries is decreased by 1 each time you guess wrong")
			fmt.Println("4. Number of tries is decreased by 1 each time you guess right")
			fmt.Print("\n\n")
			fmt.Print("Enter Any Button To Navigate Back\n")

			var c string
			fmt.Scanln(&c)
			helpers.ClearScreen()

		} else if choice == 3 {
			helpers.ClearScreen()
			fmt.Println("Game Exit")
			break
		} else {
			helpers.ClearScreen()
			continue
		}
	}

}
