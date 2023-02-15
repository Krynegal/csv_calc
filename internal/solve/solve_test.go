package solve

import (
	"csvCalc/internal/parsers"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func toRecords(s string) ([][]string, error) {
	r := strings.NewReader(s)
	reader := csv.NewReader(r)
	reader.Comma = ','
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func TestSolveTable(t *testing.T) {
	tests := []struct {
		name     string
		tableIn  string
		tableOut string
		err      error
	}{
		{
			name: "1",
			tableIn: "" +
				",A,B,Cell\n" +
				"1,=B2+B30,0,9\n" +
				"2,2,=Cell1+Cell30,0\n" +
				"30,0,=A2+A30,5",
			tableOut: "" +
				",A,B,Cell\n" +
				"1,16,0,9\n" +
				"2,2,14,0\n" +
				"30,0,2,5",
			err: nil,
		},
		{
			name: "2",
			tableIn: "" +
				",A,B,Cell\n" +
				"1,1,0,1\n" +
				"2,2,=B30+A1,0\n" +
				"30,0,=B2+A1,5",
			tableOut: "",
			err:      errors.New("cannot solve this table"),
		},
		{
			name: "3",
			tableIn: "" +
				",A,B,Cell\n" +
				"1,=Cell1+B1,0,9\n" +
				"2,2,=A2+A1,0\n" +
				"30,0,=B2+A1,5",
			tableOut: "" +
				",A,B,Cell\n" +
				"1,9,0,9\n" +
				"2,2,11,0\n" +
				"30,0,20,5",
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recordsIn, err := toRecords(tt.tableIn)
			if err != nil {
				log.Fatal(err)
			}
			tableIn, calcCell, err := parsers.ParseRecords(recordsIn)
			if err != nil {
				fmt.Println(err)
			}
			if err = SolveTable(tableIn, calcCell); err != nil {
				assert.Error(t, err, tt.err)
			}

			if tt.tableOut != "" {
				recordsOut, err := toRecords(tt.tableOut)
				if err != nil {
					log.Println(err)
				}
				tableOut, _, err := parsers.ParseRecords(recordsOut)
				if err != nil {
					fmt.Println(err)
				}
				assert.Equal(t, tableIn, tableOut)
			}
		})
	}
}
