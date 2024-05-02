package gophercises

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("quiz.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	// CSV Reader
	reader := csv.NewReader(file)

	//Enumerate CSV records
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println(record)
	}
}
