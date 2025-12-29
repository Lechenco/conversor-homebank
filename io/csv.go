package io

import (
	"encoding/csv"
	"os"
)

func ReadCSV(input string) [][]string {
	f, err := os.Open(input)
	if err != nil {
		panic("Unable to read input file " + input + ": " + err.Error())
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic("Unable to parse file as CSV " + input + ": " + err.Error())
	}

	return records
}
