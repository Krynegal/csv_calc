package main

import (
	"csvCalc/internal/parsers"
	"csvCalc/internal/solve"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]
	records, err := parsers.ReadCSV(filePath)
	if err != nil {
		log.Fatal("error while reading file")
	}
	table, calcCells, err := parsers.ParseRecords(records)
	if err != nil {
		log.Fatal(err)
	}
	if len(calcCells) != 0 {
		if err = solve.SolveTable(table, calcCells); err != nil {
			log.Fatal(err)
		}
	}
	parsers.ParseToCSV(table)
}
