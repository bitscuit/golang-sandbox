package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(printMessage())
	readCSV()
}

func printMessage() string {
	return "test"
}

func readCSV() {
	// in := `question,answer`
	file, _ := os.Open("problems.csv")
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[0], ",", record[1])
	}
}
