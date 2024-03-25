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

	// totalQuestionNumber := 0
	// correctAnswers := 0

	// arithmeticOperationMap := make(map[rune]func(leftOperand int , rightOperand int ) int{
	// 	'+': func(leftOperand int , rightOperand int ) int {
	// 		return leftOperand + rightOperand
	// 	},
	// 	'-' : func(leftOperand int , rightOperand int ) int {
	// 		return leftOperand - rightOperand
	// 	},
	// 	'*' : func(leftOperand int , rightOperand int ) int {
	// 		return leftOperand * rightOperand
	// 	},
	// 	'/' : func(leftOperand int , rightOperand int ) int {
	// 		return leftOperand / rightOperand
	// 	},
	// })

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

		leftOperand , isLeftParseError := strconv.ParseInt(operands[0] , 10 , 16)
		rightOperand , isRightParseError := strconv.ParseInt(operands[1] , 10 , 16)
		result ,  isResultParseError :=  strconv.ParseInt(data[1] , 10 , 16) 

		if isLeftParseError != nil || isRightParseError != nil || isResultParseError != nil {
			log.Fatal("Not Valid Number")
		}

		fmt.Printf("lhs : %v , rhs : %v , total : %v \n" , leftOperand , rightOperand , leftOperand + rightOperand)

	}

	defer file.Close()

}
