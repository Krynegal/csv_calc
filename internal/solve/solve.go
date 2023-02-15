package solve

import (
	"csvCalc/internal/models"
	"errors"
	"fmt"
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

func makeExpression(v1, v2, op string) (models.Expression, error) {
	var expr models.Expression
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
	return expr, nil
}

func parseExpr(table map[models.Cell]string, val string) (string, string, string) {
	var args []string
	for _, op := range []string{"+", "-", "*", "/"} {
		if strings.Index(val[1:], op) != -1 {
			args = strings.Split(val[1:], op)

			cell1 := strToCell(args[0])
			cell2 := strToCell(args[1])

			v1 := table[cell1]
			v2 := table[cell2]
			return v1, v2, op
		} else {
			continue
		}
	}
	return "", "", ""
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

func SolveTable(table map[models.Cell]string, cellsToCalc []models.Cell) error {
	i := 0
	lastLen := len(cellsToCalc)
	for len(cellsToCalc) != 0 {
		if i == lastLen {
			return errors.New("this table cannot be solved")
		}
		val := table[cellsToCalc[i]]
		operand1, operand2, op := parseExpr(table, val)
		if op == "/" && operand2 == "0" {
			return errors.New("division by zero")
		}
		if operand1 == val {
			return fmt.Errorf("cycle: %v %v", val, operand1)
		} else if operand2 == val {
			return fmt.Errorf("cycle: %v %v", val, operand2)
		}
		if strings.Contains(operand1, "=") || strings.Contains(operand2, "=") {
			i++
			continue
		}

		expression, err := makeExpression(operand1, operand2, op)
		if err != nil {
			return err
		}
		result := calcExpression(expression)
		table[cellsToCalc[i]] = result
		cellsToCalc = append(cellsToCalc[:i], cellsToCalc[i+1:]...)
		if i != 0 {
			i--
		}
	}
	return nil
}
