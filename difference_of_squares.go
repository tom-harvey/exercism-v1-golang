// Package diffsquares makes semi-pointless calculations.
// Note: there is some overflow avoidance but no overflow detection
package diffsquares

import (
	"math"
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
)

var maxSumN, maxFaulhaberN, maxSquareN int

func init() {
	// this could be a constant if we knew the size of int
	maxFaulhaberN = int(math.Pow(float64(maxInt)/2.0, 1.0/3.0))
}

// SumOfSquares calculates the sum from 1 to n of n*n
// It returns zero for n less than 1.
func SumOfSquares(n int) int {
	var s int
	// Faulhaber's formula for power 2: (n(n+1)(2n+1))/6
	if n > 0 {
		if n <= maxFaulhaberN {
			s = n * (n + 1) * (2*n + 1) / 6
		} else { // TODO overflow-related code is untested!
			s = SumOfSquares(maxFaulhaberN)
			for i := maxFaulhaberN + 1; i <= n; i++ {
				s += i * i
			}
		}
	}
	return s
}

// SquareOfSums calculates the sum from 1 to n, squared.
// It returns zero for n less than 1.
func SquareOfSums(n int) int {
	var result int
	if n > 0 {
		// Faulhaber's formula for power 1:(n(n+1))/2
		// modified to avoid unnecessary overflow
		odd := n & 1
		sum := (n-odd+1)*(n>>1) + odd*n
		result = sum * sum
	}
	return result
}

// Difference returns the difference between SquareOfSums and SumOfSquares
// It returns -1 in case of overflow
func Difference(n int) int {
	// the answer is actually n(3n³ + 2n² - 3n - 2)/12 but that will overflow
	// sooner than:
	return SquareOfSums(n) - SumOfSquares(n)
}
