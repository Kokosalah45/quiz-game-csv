package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	DEFAULT_PROBLEMSET = "problems"
)

func main() {

	problemSet := DEFAULT_PROBLEMSET

	if len(os.Args) > 1 {
		problemSet = os.Args[1]
	}
	path := fmt.Sprintf("questions-repo/%v.csv", problemSet)

	file, err := os.Open(path)

	if err != nil {
		log.Fatal("No File found")
	}

	reader := csv.NewReader(file)

	totalQuestionNumber := 0
	correctAnswers := 0

	for {
		data, err := reader.Read()

		if err == io.EOF {
			break
		}
		operation := data[0]

		r := regexp.MustCompile(`(\\|\+|-|\*)`)

		operatorIndex := r.FindStringIndex(operation)[0]

		operator := string(operation[operatorIndex])

		operands := strings.Split(operation, operator)

		fmt.Fprintf(os.Stdout, "What is the result of that equation ?\n")
		fmt.Fprintf(os.Stdout, "%s %s %s = ?\n", operands[0], operator, operands[1])
		fmt.Fprintf(os.Stdout, "Type your result here : ")
		userInput := 0
		_, isScanError := fmt.Fscan(os.Stdin, &userInput)

		if isScanError != nil {
			log.Fatal("Wrong entry")

		}

		result, isParseError := strconv.ParseInt(data[1], 10, 0)

		if  isParseError != nil {
			log.Fatal("Not Valid Number")
		}

		if(int(result) == userInput){
			correctAnswers++
		}
		totalQuestionNumber++


	}
	fmt.Fprintf(os.Stdout, "You answered %v correct out of %v question" , correctAnswers, totalQuestionNumber)


	defer file.Close()

}
