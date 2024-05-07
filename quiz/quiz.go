package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	var shuffle bool

	flag.BoolVar(&shuffle, "r", false, "Randomize questions")
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s", *csvFilename)
		os.Exit(1)
	}
	defer file.Close()

	// CSV Reader
	reader := csv.NewReader(file)

	//Enumerate CSV records
	csv_records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Shuffle questions if necessary
	if shuffle {
		rand.New(rand.NewSource(time.Now().Unix()))
		for i := len(csv_records) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			csv_records[i], csv_records[j] = csv_records[j], csv_records[i]
		}
	}

	// Run quiz loop!
	round := 0
	score := 0
	for _, record := range csv_records {
		prob := problem{
			q: record[0],
			a: record[1],
		}
		score += question(prob, round)
		round++
	}

	// Print score and exit
	fmt.Printf("You scored %d/12!\n", score)
	os.Exit(0)
}

type problem struct {
	q string
	a string
}

func question(record problem, num int) int {
	var userInput string
	fmt.Printf("Problem #%d: %s = ", num+1, record.q)
	fmt.Scanln(&userInput)
	userInput = strings.ReplaceAll(userInput, " ", "")
	if userInput == record.a {
		return 1
	}
	return 0
}
