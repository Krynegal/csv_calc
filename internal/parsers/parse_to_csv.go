package parsers

import (
	"csvCalc/internal/models"
	"encoding/csv"
	"log"
	"os"
	"sort"
)

func ParseToCSV(table map[models.Cell]string) [][]string {
	var (
		columns, rows []string
		records       [][]string
	)

	columnsMap := make(map[string]struct{})
	rowsMap := make(map[string]struct{})

	columns = append(columns, "")
	for cell := range table {
		row, column := cell.Row, cell.Column
		if _, ok := columnsMap[column]; !ok {
			columnsMap[column] = struct{}{}
			columns = append(columns, column)
		}
		if _, ok := rowsMap[row]; !ok {
			rowsMap[row] = struct{}{}
			rows = append(rows, row)
		}
	}

	sort.Strings(columns)
	sort.Strings(rows)

	records = append(records, columns)
	for _, row := range rows {
		recRow := make([]string, len(columns))
		recRow[0] = row
		ind := 1
		for _, column := range columns[1:] {
			cell := models.Cell{
				Row:    row,
				Column: column,
			}
			recRow[ind] = table[cell]
			ind++
		}
		records = append(records, recRow)
	}

	writer := csv.NewWriter(os.Stdout)
	writer.Comma = ','
	err := writer.WriteAll(records)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
