package internal

import (
	"csvCalc/internal/models"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func strToCell(val string) models.Cell {
	var col, row string
	for i, c := range val {
		if unicode.IsDigit(c) {
			col = val[:i]
			row = val[i:]
			break
		}
	}
	return models.Cell{
		Row:    row,
		Column: col,
	}
}

func parseExpr(table map[models.Cell]string, val string) (models.Expression, error) {
	var expr models.Expression
	var args []string
	for _, op := range []string{"*", "/", "+", "-"} {
		if strings.Index(val[1:], op) != -1 {
			args = strings.Split(val[1:], op)

			cell1 := strToCell(args[0])
			cell2 := strToCell(args[1])

			v1 := table[cell1]
			v2 := table[cell2]

			arg1, err := strconv.Atoi(v1)
			if err != nil {
				return expr, err
			}
			arg2, err := strconv.Atoi(v2)
			if err != nil {
				return expr, err
			}
			expr = models.Expression{
				Arg1:     arg1,
				Arg2:     arg2,
				Operator: op,
			}
			break
		} else {
			continue
		}
	}
	return expr, nil
}

func calcExpression(expr models.Expression) string {
	var result int
	switch expr.Operator {
	case "+":
		result = expr.Arg1 + expr.Arg2
	case "-":
		result = expr.Arg1 - expr.Arg2
	case "*":
		result = expr.Arg1 * expr.Arg2
	case "/":
		result = expr.Arg1 / expr.Arg2
	}
	return strconv.Itoa(result)
}

func SolveTable(table map[models.Cell]string, calcCells []models.Cell) {
	for _, calcCell := range calcCells {
		expression, err := parseExpr(table, table[calcCell])
		if err != nil {
			log.Fatal("solve table error")
		}
		result := calcExpression(expression)
		table[calcCell] = result
	}
}
