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
		quiz(record, counter)
		counter++
	}
}

func quiz(record []string, num int) {
	if len(record) >= 2 {
		question := record[0]
		answer := record[1]
		fmt.Println("Problem: ", question)
	}
}
