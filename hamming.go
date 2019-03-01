// Package hamming has Just One Job
package hamming

import (
	"errors"
)

// Distance computes the Hamming distance between two strings (as bytes)
func Distance(a, b string) (int, error) {
	d := 0
	err := error(nil)
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] { // NOTE: comparing as bytes, not as runes
				d++
			}
		}
	} else {
		d = -1 // don't let callers ignore err
		err = errors.New("lengths unequal")
	}
	return d, err
}
