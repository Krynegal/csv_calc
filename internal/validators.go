package internal

import (
	"strconv"
	"strings"
)

func CheckColumns(columns []string) bool {
	c := make(map[string]struct{}, 0)
	for _, column := range columns {
		if _, ok := c[column]; ok {
			return false
		}
		c[column] = struct{}{}
	}
	return true
}

func IsExpression(val string) bool {
	if len(val) < 4 || val[0] != '=' || !strings.ContainsAny(val, "+-*/") {
		return false
	}
	return true
}

func IsNumber(val string) bool {
	_, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return true
}
