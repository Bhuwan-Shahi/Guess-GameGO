//GREET THE USER
//SELECT THE DIFFICULTY LEVEL OF THE USER
//ASK FOR THE INPUT
//UPDATE THE ANSWER
//CHECK WITH THE CURRENT ANSWER
//SET THE QUITING CONDITION OR EDGES CASES OF WHEN THE APPLICATION IS ENDED

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var difficultyLeevel int
var answer int

func greetUser() {
	fmt.Println("Welcome to Number Guessing Game!!")
}

func chooseLevel() {
	fmt.Println("We have 3 level of game")
	fmt.Println("1. Easy(10 Chances)")
	fmt.Println("2. Medium(5 Chances)")
	fmt.Println("3. Hard(3 Chances)")
	fmt.Println("Enter which level you want:")
	fmt.Scanln(&difficultyLeevel)

}

func generateComputerAnswer() {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	answer = generator.Intn(100)
	fmt.Println("________________________________________")
	fmt.Println("I'm thinking the number between 0 - 100")
	fmt.Println("________________________________________")

}

func checkWin(noofChances int) {
	var guess int
	chance := 0

	for chance < noofChances {
		fmt.Println("Enter your Guess:")
		fmt.Scanln(&guess)
		if guess == answer {
			fmt.Println("You win!!!!!!")
			fmt.Println("You have guessed the correct answe in :", guess, "guesses")
			break
		} else if guess < answer {
			fmt.Println("The answer is greater than :", guess)
		} else {
			fmt.Println("the anser is less than :", guess)
		}
		chance++
	}
	if guess != answer {
		fmt.Println("________________________________________")
		fmt.Println("You loose the no of chances exceedes")
		fmt.Println("The answer was :", answer)
		fmt.Println("________________________________________")
	}

}

func playGame() {
	switch difficultyLeevel {
	case 1:
		fmt.Print("You have Choosen Easy mode.")
		checkWin(10)
	case 2:
		fmt.Println("You have Choosen Intermediate mode")
		checkWin(5)
	case 3:
		fmt.Println("You have Choosen Hard mode")
		checkWin(3)
	default:
		fmt.Println("You have Choosen the level we don't have!!")
		chooseLevel()
	}
}
