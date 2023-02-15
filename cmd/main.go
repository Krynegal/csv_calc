package main

import (
	"csvCalc/internal"
	"csvCalc/internal/parsers"
	"fmt"
	"log"
)

func main() {
	filePath := "file.csv"
	records, err := parsers.ReadCSV(filePath)
	if err != nil {
		log.Fatal("error while reading file")
	}
	table, calcCell := parsers.ParseRecords(records)
	fmt.Println(table)
	fmt.Println(calcCell)

	internal.SolveTable(table, calcCell)
	fmt.Println(table)
	//t := map[parsers.Cell]string{
	//	parsers.Cell{"1", "A"}:     "1",
	//	parsers.Cell{"1", "B"}:     "0",
	//	parsers.Cell{"1", "Cell"}:  "1",
	//	parsers.Cell{"2", "A"}:     "2",
	//	parsers.Cell{"2", "B"}:     "6",
	//	parsers.Cell{"2", "Cell"}:  "0",
	//	parsers.Cell{"30", "A"}:    "0",
	//	parsers.Cell{"30", "B"}:    "1",
	//	parsers.Cell{"30", "Cell"}: "5",
	//}

	parsers.ParseToCSV(table)
}
