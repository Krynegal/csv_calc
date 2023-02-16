package internal

import (
	"strings"
	"unicode"
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
	if len(val) < 6 || val[0] != '=' || !strings.ContainsAny(val, "+-*/") {
		return false
	}
	return true
}

func IsNumber(val string) bool {
	for _, c := range val {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
