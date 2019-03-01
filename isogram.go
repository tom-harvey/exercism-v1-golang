// Package tells us if a string is an Isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram takes a string and returns if its an isogram or not.
func IsIsogram(input string) bool {
	var results = make(map[string]int)
	for i := 0; i < len(input); i++ {
		char := strings.ToUpper(input[i])
		isLetter := unicode.IsLetter(rune(input[i]))
		if isLetter && results[char] != 0 {
			results[char] += 1
		} else if isLetter {
			results[char] = 1
		}

		for _, v := range results {
			if v > 1 {
				return false
			}
		}
	}
	return true
}
