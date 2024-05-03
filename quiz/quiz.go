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
	score := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		if len(record) >= 2 {
			score += question(record, counter)
		} else {
			break
		}
		counter++
	}
	fmt.Printf("You scored %d/12!\n", score)
}

func question(record []string, num int) int {
	var userInput string
	question := record[0]
	answer := record[1]
	fmt.Printf("Problem #%d: %s = ", num+1, question)
	fmt.Scanln(&userInput)
	if userInput == answer {
		return 1
	}
	return 0
}
