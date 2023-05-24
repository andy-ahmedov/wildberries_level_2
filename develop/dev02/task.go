package string_unpacker

import (
	"fmt"
	"strconv"
	"strings"
)

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	if isDigit(s[0]) {
		err := fmt.Errorf("invalid string")
		return "", err
	}
	
	lastChar := rune(s[0])

	var result strings.Builder
	var count string
	pad := " "
	s = (s + pad)[1:]

	for _, ch := range s {
		if isDigit(ch) {
			count += string(ch)
		} else {
			if len(count) > 0 {
				c, _ := strconv.Atoi(count)
				for j := 0; j < c; j++ {
					result.WriteRune(lastChar)
				}
			} else {
				result.WriteRune(lastChar)
			}
			count = ""
			lastChar = ch
		}
	}

	return result.String(), nil
}

func isDigit[T rune | byte](ch T) bool {
	return ch >= '0' && ch <= '9'
}