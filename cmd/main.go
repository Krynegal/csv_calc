package main

import (
	"csvCalc/internal/parsers"
	"csvCalc/internal/solve"
	"fmt"
	"log"
)

func main() {
	filePath := "file.csv"
	records, err := parsers.ReadCSV(filePath)
	if err != nil {
		log.Fatal("error while reading file")
	}
	table, calcCells, err := parsers.ParseRecords(records)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(table)
	fmt.Println(calcCells)

	if len(calcCells) != 0 {
		if err = solve.SolveTable(table, calcCells); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(table)
	parsers.ParseToCSV(table)
}
