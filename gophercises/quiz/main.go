package main

import (
	"fmt"
)

func main() {

	qaList := readFromFile()

	correctAnswers := 0
	totalQuestions := len(qaList)

	for _, q := range qaList {
		fmt.Println(q.question)
		var input string
		fmt.Scanln(&input)
		if input == q.answer {
			fmt.Println("Correct")
			correctAnswers = correctAnswers + 1
		} else {
			fmt.Println("Not correct bro")
		}
	}

	var passPcnt float64 = float64(correctAnswers) * 100 / float64(totalQuestions)
	fmt.Println("Total questions: ", totalQuestions)
	fmt.Println("Total correct: ", correctAnswers)
	fmt.Println("Percentage correct: ", passPcnt)

}
