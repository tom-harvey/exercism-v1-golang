// Package acronym generates acronyms
package acronym

import (
	//"strings"      // used in SLOWER version
	"unicode"
	//"unicode/utf8" // used in SLOWER version
)

// Abbreviate returns the acronym of the phrase in s
func Abbreviate(s string) string {
	var initials []rune
	/* SLOWER: 3500 ns/op
	for _, word := range strings.Fields(s) {
		for _, hword := range strings.Split(word, "-") {
			r, _ := utf8.DecodeRuneInString(hword)
			initials = append(initials, unicode.ToUpper(r))
		}
	} */
	// FASTER: 2200 ns/op because it walks through string only once
	start := true
	for _, r := range s {
		if start && unicode.IsLetter(r) {
			initials = append(initials, unicode.ToUpper(r))
			start = false
		} else if unicode.IsSpace(r) || r == '-' {
			start = true
		}
	}
	return string(initials)
}
