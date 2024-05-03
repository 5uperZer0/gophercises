package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	// CSV Reader
	reader := csv.NewReader(file)

	//Enumerate CSV records
	counter := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		question(record, counter)
		counter++
	}
}

func question(record []string, num int) bool {
	if len(record) >= 2 {
		var userInput string
		question := record[0]
		answer := record[1]
		fmt.Printf("Problem #", num+1, ": ", question, " = ")
		fmt.Scan(&userInput)
		if userInput == answer {
			return true
		}
		return false
	}
}
