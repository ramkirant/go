package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
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

	r := csv.NewReader(strings.NewReader(fileContent))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if len(record) != 0 {
			q := quizQAndA{}
			q.question = record[0]
			q.answer = record[1]
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
