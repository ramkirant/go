package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	qaList := readFromFile()

	sleepSeconds := 20
	correctAnswers := 0
	c := make(chan string)

	fmt.Println("Welcome to the fun quiz. You have ", sleepSeconds, " seconds to finish the quiz")
	fmt.Println("Press any key to start the quiz")
	fmt.Scanln()

	go executeTheQuiz(qaList, &correctAnswers, c)

	go sleep(sleepSeconds, c)

	for l := range c {
		if l == "Kill" {
			fmt.Println("Quiz ended successfully")
		} else {
			fmt.Println("Quiz failed")
		}

		totalQuestions := len(qaList)
		var passPcnt float64 = float64(correctAnswers) * 100 / float64(totalQuestions)
		fmt.Println("Total questions: ", totalQuestions)
		fmt.Println("Total correct: ", correctAnswers)
		fmt.Println("Percentage correct: ", passPcnt)
		os.Exit(0)
	}
}

func sleep(sleepSeconds int, c chan string) {
	time.Sleep(time.Duration(sleepSeconds * int(time.Second)))
	fmt.Println("\nTimeout. Sorry bro")
	c <- "Kill"
}

func executeTheQuiz(qaList quizQAndAList, correctAnswers *int, c chan string) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range qaList {
		nP := r.Intn(len(qaList) - 1)
		qaList[i], qaList[nP] = qaList[nP], qaList[i]
	}

	for _, q := range qaList {
		fmt.Println(q.question)
		var input string
		fmt.Scanln(&input)
		if input == q.answer {
			fmt.Println("Correct")
			*correctAnswers = *correctAnswers + 1
		} else {
			fmt.Println("Not correct bro")
		}
	}

	c <- "Kill"
}
