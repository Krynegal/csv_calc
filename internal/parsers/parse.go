package parsers

import (
	"csvCalc/internal"
	"csvCalc/internal/models"
	"encoding/csv"
	"errors"
	"os"
)

func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if err != nil {
			err = closeErr
		}
	}(file)

	reader := csv.NewReader(file)
	reader.Comma = ','
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func ParseRecords(records [][]string) (map[models.Cell]string, []models.Cell, error) {
	var cellsToCalc []models.Cell
	table := make(map[models.Cell]string, 0)
	prevRowName := ""
	columnName := records[0]
	if !internal.CheckColumns(records[0]) {
		return nil, nil, errors.New("wrong columns format")
	}
	for _, record := range records[1:] {
		row := record
		rowName := row[0]
		if prevRowName == rowName {
			return nil, nil, errors.New("wrong table format")
		}
		prevRowName = rowName
		for ind, val := range row[1:] {
			cell := models.Cell{
				Column: columnName[ind+1],
				Row:    rowName,
			}
			if internal.IsExpression(val) {
				cellsToCalc = append(cellsToCalc, cell)
			} else if !internal.IsNumber(val) {
				return nil, nil, errors.New("only numbers and expressions can be cell's value")
			}
			table[cell] = val
		}
	}
	return table, cellsToCalc, nil
}
