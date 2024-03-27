package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
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


func runQuiz(problemSet string, quizChan chan Quiz) {


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

func terminateAppWorker(timer time.Timer, solved chan Quiz , wg *sync.WaitGroup)  {
	select {
	case <-timer.C:
		fmt.Fprintf(os.Stdout, "\nTime is up!\n")
		wg.Done()
	case data := <-solved:
		wg.Done()
		fmt.Fprintf(os.Stdout, "You answered %v correct out of %v questions\n", data.correctAnswers, data.totalQuestionNumber)
	}
}

func main() {

	problemSetPtr := flag.String("source", DEFAULT_PROBLEMSET , "source of questions")
	timeLimitPtr  := flag.Duration("time", DEFAULT_TIME, "time limit for the quiz")

	flag.Parse()

	problemSet := *problemSetPtr
	timeLimit := *timeLimitPtr	

	var wg sync.WaitGroup


	quizChan := make(chan Quiz)
	timer := time.NewTimer(timeLimit)

	wg.Add(1)
	go runQuiz(problemSet, quizChan)
	go terminateAppWorker(*timer, quizChan , &wg)
	wg.Wait()



}
