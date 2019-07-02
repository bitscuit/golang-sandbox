package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Quiz Game")
	records := readCSV()
	fmt.Println(records) // DEBUG
	fmt.Println(len(records)) // DEBUG
	
}

func readCSV() [][]string {
	// in := `question,answer`
	file, _ := os.Open("problems.csv")
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}
