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

	// Process flags
	var shuffle bool
	flag.BoolVar(&shuffle, "r", false, "Randomize questions")
	timeLimit := flag.Int("limit", 3, "the time limit for the quiz in seconds")
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

	// Instantiate quiz timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	defer timer.Stop()

	// Run quiz loop!
	round := 0
	score := 0

	for _, record := range csv_records {

		fmt.Printf("Problem #%d: %s = ", round+1, record[0])
		answer_ch := make(chan string)
		go func() {
			var in string
			fmt.Scanln(&in)
			in = strings.ReplaceAll(in, " ", "")
			answer_ch <- in
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d/12!\n", score)
			return
		case answer := <-answer_ch:
			if answer == record[1] {
				fmt.Println("Correct!")
				score++
			}
		}
	}

	// Print score and exit
	fmt.Printf("\nYou scored %d/12!\n", score)
	os.Exit(0)
}
