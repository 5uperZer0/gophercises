package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	var shuffle bool

	flag.BoolVar(&shuffle, "r", false, "Randomize questions")
	flag.Parse()

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	// CSV Reader
	reader := csv.NewReader(file)

	//Enumerate CSV records
	var csv_records [][]string
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
			csv_records = append(csv_records, record)
		} else {
			continue
		}
	}
	if shuffle {
		rand.New(rand.NewSource(time.Now().Unix()))
		for i := len(csv_records) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			csv_records[i], csv_records[j] = csv_records[j], csv_records[i]
		}
		for _, record := range csv_records {
			score += question(record, counter)
			counter++
		}
	} else {
		for _, record := range csv_records {
			score += question(record, counter)
			counter++
		}
	}

	fmt.Printf("You scored %d/12!\n", score)
}

func question(record []string, num int) int {
	var userInput string
	question := record[0]
	answer := record[1]
	fmt.Printf("Problem #%d: %s = ", num+1, question)
	fmt.Scanln(&userInput)
	userInput = strings.ReplaceAll(userInput, " ", "")
	if userInput == answer {
		return 1
	}
	return 0
}
