package main

import (
	"testing"
)

func TestReadCSV_ShouldReturnAllQuestionAnswerPairs(t *testing.T) {
	expected := 3 // TODO: make this easily changeable
	records := readCSV()
	actual := len(records)
	if expected != actual {
		t.Errorf("readCSV == %d, expected == %d", actual, expected)
	}
}

func TestGame_ShouldGiveCorrectScore(t *testing.T) {

}

func TestAskQuestion(t *testing.T) {

}

func TestAnswerQuestion(t *testing.T) {
	
}
