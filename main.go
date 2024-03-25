package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

const DEFAULT_PROBLEMSET = "problems"

func main() {

	problemSet := DEFAULT_PROBLEMSET

	if(len(os.Args) > 1){
		problemSet = os.Args[1]
	}

	file , err := os.Open(fmt.Sprintf("questions-repo/%v.csv" , problemSet))

	if err != nil {
		log.Fatal("No File found")
	}

	reader := csv.NewReader(file)
	
	  
	for {
		data , err := reader.Read()
		if (err == io.EOF){
			break
		}
		fmt.Println(data)

	}

	defer file.Close()


}