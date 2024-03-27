package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	DEFAULT_PROBLEMSET = "problems"
	DEFAULT_TIME = 5
	DEFAULT_TIME_UNIT = "s"
)

type Quiz struct {
	correctAnswers int
	totalQuestionNumber int
}

func getCorrespondingTimeUnit(timeUnit string) time.Duration {
	switch timeUnit {
	case "s":
		return time.Second
	case "m":
		return time.Minute
	case "h":
		return time.Hour
	default:
		return time.Second
	}
}

func runQuiz(timer time.Timer, done chan Quiz)  {
	select {
	case <-timer.C:
		fmt.Fprintf(os.Stdout, "\nTime is up!\n")
		os.Exit(0)
	case quiz := <-done:
		fmt.Fprintf(os.Stdout, "You answered %v correct out of %v questions\n", quiz.correctAnswers, quiz.totalQuestionNumber)
	}
}

func main() {

	var timeLimit time.Duration = DEFAULT_TIME
	timeUnit := DEFAULT_TIME_UNIT
	

	probmeSetPtr := flag.String("source", DEFAULT_PROBLEMSET , "source of questions")
	timeLimitPtr := flag.Int("time", DEFAULT_TIME, "time limit for the quiz")
	timeUnitPtr := flag.String("time-unit", DEFAULT_TIME_UNIT, "time unit for the quiz")

	flag.Parse()

	problemSet := *probmeSetPtr
	timeLimit = time.Duration(*timeLimitPtr)
	timeUnit = *timeUnitPtr

	path := fmt.Sprintf("questions-repo/%v.csv", problemSet)

	file, err := os.Open(path)

	quizChan := make(chan Quiz)
	timer := time.NewTimer(timeLimit * getCorrespondingTimeUnit(timeUnit))

	go runQuiz(*timer, quizChan)



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
		fmt.Fprintf(os.Stdout, "What is the result of that equation ?\n")
		fmt.Fprintf(os.Stdout, "%s = ?\n", operation)
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
	quiz := Quiz{correctAnswers, totalQuestionNumber}
	quizChan <- quiz;

	defer close(quizChan)
	defer file.Close()

}
