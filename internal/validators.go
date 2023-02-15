package internal

import "unicode"

func ValidateNum(str string) bool {
	for _, symb := range str {
		if !unicode.IsDigit(symb) {
			return false
		}
	}

	return true
}

func ValidateValue(str string) (string, bool) {
	strRune := []rune(str)

	if len(str) != 0 && (!unicode.IsGraphic(strRune[0]) || strRune[0] == '=') {
		return "string", true
	}

	if ValidateNum(str) {
		return "int", true
	}

	return "", false
}
