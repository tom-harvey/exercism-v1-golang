// Package luhn validates a luhn checksum
package luhn

// Package luhn validates a luhn checksum

var dbl2 = [...]int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf,
	0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// Valid returns true if the last digit in a string containing at least two
// ASCII digits and optionally ASCII spaces, and no other characters, is a
// valid luhn check digit.
func Valid(s string) bool {
	var dlen, sum int
	check := -1

	for i := len(s) - 1; i >= 0; i-- {
		if ch := int(s[i]); ch != ' ' {
			ch -= '0'
			if ch < 0 || ch > 9 {
				return false // I don't like mid-function returns, but...
			} else if check < 0 {
				check = ch
			} else {
				dlen++
				sum += dbl2[ch|dlen&1<<4]
			}
		}
	}
	return dlen > 0 && (sum*9)%10 == check
}

//
// ----- a version that is slightly slower on my machine
//
//                 0  1  2  3  4  5  6  7  8  9
var dbl = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// ValidSlightlySlower is like Valid() above, but slightly slower.
func ValidSlightlySlower(s string) bool {
	var dlen, sum int
	check := -1

	for i := len(s) - 1; i >= 0; i-- {
		if ch := int(s[i]); ch != ' ' {
			ch -= '0'
			if ch < 0 || ch > 9 {
				return false // I don't like mid-function returns, but...
			} else if check < 0 {
				check = ch
			} else {
				dlen++
				if dlen&1 == 0 {
					sum += ch
				} else {
					sum += dbl[ch]
				}
			}
		}
	}
	return dlen > 0 && (sum*9)%10 == check
}
