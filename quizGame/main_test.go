package main

import (
	"testing"
	"encoding/csv"
)

func TestPrintMessage(t *testing.T) {
	var got = printMessage()
	var want = "test"
	if got != want {
		t.Errorf("printMessage() == %s, want %s", got, want)
	}
}

func TestReadProblemsCSV(t *testing.T) {
	
}
