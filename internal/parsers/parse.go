package parsers

import (
	"csvCalc/internal"
	"csvCalc/internal/models"
	"encoding/csv"
	"fmt"
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

func ParseRecords(records [][]string) (map[models.Cell]string, []models.Cell) {
	var equations []models.Cell
	table := make(map[models.Cell]string, 0)
	//prevRowName := ""
	columnName := records[0]
	for _, record := range records[1:] {
		row := record
		rowName := row[0]
		//if prevRowName == rowName {
		//	fmt.Println("here")
		//} else {
		//	prevRowName = rowName
		//}
		for ind, val := range row[1:] {
			valType, ok := internal.ValidateValue(val)
			if !ok {
				fmt.Println('1')
			}

			cell := models.Cell{
				Column: columnName[ind+1],
				Row:    rowName,
			}

			if valType == "string" {
				//equations = append(equations, columnName[ind+1]+rowName)
				equations = append(equations, cell)
			}

			//table[columnName[ind+1]+rowName] = val
			table[cell] = val
		}
	}
	return table, equations
}
