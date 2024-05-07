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

type problem struct {
	q string
	a string
}

func main() {

	// Process flags
	var shuffle bool
	flag.BoolVar(&shuffle, "r", false, "Randomize questions")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'question, answer'")
	flag.Parse()

	// Import CSV data and transform to problem struct
	problems, err := formatCSV(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s", *csvFilename)
		os.Exit(1)
	}

	// Shuffle questions
	if shuffle {
		rand.New(rand.NewSource(time.Now().Unix()))
		for i := len(problems) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			problems[i], problems[j] = problems[j], problems[i]
		}
	}

	score := quiz(problems, *timeLimit)

	// Print score and exit
	fmt.Printf("You scored %d/12\n", score)
	os.Exit(0)
}

func formatCSV(filename string) ([]problem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	csv_records, err := reader.ReadAll()
	ret := make([]problem, len(csv_records))

	for i, line := range csv_records {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret, err
}

func quiz(problems []problem, timeLimit int) int {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	defer timer.Stop()

	score := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)
		answer := getUserInput(timer)
		if answer == problem.a {
			fmt.Println("Correct!")
			score++
		} else if answer == "END" {
			break
		}
	}
	return score
}

func getUserInput(timer *time.Timer) string {
	answer_ch := make(chan string)
	go func() {
		var in string
		fmt.Scanln(&in)
		in = strings.ReplaceAll(in, " ", "")
		answer_ch <- in
	}()

	select {
	case <-timer.C:
		fmt.Println("\nTime's up!")
		return "END"
	case answer := <-answer_ch:
		return answer
	}
}
