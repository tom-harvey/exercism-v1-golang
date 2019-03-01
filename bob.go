// Package bob responds as Bob
package bob

import (
	"strings"
	"unicode"
)

// Hey returns an appropriate remark, for Bob
func Hey(remark string) string {
	r := "Whatever."
	hasPrint := false
	if len(remark) > 0 {
		isQuestion := strings.HasSuffix(strings.TrimSpace(remark), "?")
		// yelling is some uppercase and no lower case
		hasLower := false
		hasUpper := false
		for _, c := range remark {
			if unicode.IsLower(c) {
				hasLower = true
			} else if unicode.IsUpper(c) {
				hasUpper = true
			}
			if !unicode.IsSpace(c) {
				hasPrint = true
			}
		}
		isYell := hasUpper && !hasLower
		if isQuestion {
			if isYell {
				r = "Calm down, I know what I'm doing!"
			} else {
				r = "Sure."
			}
		} else if isYell {
			r = "Whoa, chill out!"
		}
	}
	if !hasPrint {
		r = "Fine. Be that way!"
	}
	return r
}
