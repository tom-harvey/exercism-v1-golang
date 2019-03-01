// Package robotname generates unique names
package robotname

import (
	"fmt"
)

const (
	namePattern = "AA###"
	prime       = 675979
	last        = 676000 // 26*26*10*10*10
)

type Robot struct {
	name string
}

var next uint

func init() {
	testForDuplicates()
}

// Name returns the current name of a Robot
func (r *Robot) Name() string {
	if len(r.name) == 0 {
		r.name = genName()
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}

// permuteQPR permutes a number in the range [0, prime).
// Hastily and non-rigorously adapted from:
// http://preshing.com/20121224/
// how-to-generate-a-sequence-of-unique-random-integers/
func permuteQPR(x uint) uint {
	result := x
	X := uint64(x)
	if x < prime {
		result = uint((X * X) % prime)
		if x > prime/2 {
			result = prime - result
		}
	}
	return result
}

func genName() string {
	if next >= last {
		panic("no moar robotz!")
	}
	p := permuteQPR(permuteQPR(next))
	next++
	n := make([]byte, len(namePattern))
	for i, c := range namePattern {
		switch c {
		case 'A':
			n[i] = byte('A' + p%26)
			p /= 26
		case '#':
			n[i] = byte('0' + p%10)
			p /= 10
		default:
			panic("bad code in namePattern")
		}
	}
	return string(n)
}

// testForDuplicates does just that
func testForDuplicates() {
	issuedNames := make(map[string]bool)
	for i := 0; i < last; i++ {
		s := genName()
		if issuedNames[s] {
			fmt.Printf("Name #%d %s is a DUPLICATE!\n", i, s)
			panic("")
		}
		issuedNames[s] = true
	}
	next = 0
}
