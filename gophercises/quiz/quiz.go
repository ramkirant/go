package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type quizQAndA struct {
	question string
	answer   string
}

type quizQAndAList []quizQAndA

func (q *quizQAndA) convertToQ(line string) {
	elements := strings.Split(line, ",")

	if len(elements) == 2 {
		q.question = elements[0]
		q.answer = elements[1]
	}
}

func (ql *quizQAndAList) Write(bs []byte) (int, error) {
	fileContent := string(bs)
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		if len(line) != 0 {
			q := quizQAndA{}

			if line[len(line)-1] == 13 {
				line = line[0 : len(line)-1]
			}
			q.convertToQ(line)

			*ql = append(*ql, q)
		}
	}

	return len(bs), nil
}

func readFromFile() quizQAndAList {
	qaList := quizQAndAList{}

	file, err := os.Open("quiz.csv")

	if err != nil {
		fmt.Println("Error opening the file: ", err)
	}

	io.Copy(&qaList, file)

	return qaList
}
