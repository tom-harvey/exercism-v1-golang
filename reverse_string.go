// Package reverse does what it does.
package reverse

// String --as is obvious from its name--reverses a string.
// Valid utf-8 runes are reversed as runes.
//
// This routine will not reverse n-rune unnormalized unicode characters
// nicely.
//
// The fastest version is about twice as fast on my machine as the slowest,
// but is it worth the added complexity? Depends...
//
// Move runes faster without decode/encode.
// This version preserves invalid utf-8 bytes.
func String(s string) string {
	s0 := 0
	s1 := 0
	r0 := len(s)
	reverse := make([]byte, r0, r0)
	for s1 = range s {
		// move the previous encoded rune, if any
		rlen := s1 - s0
		r0 -= rlen
		r := r0
		switch rlen {
		case 4:
			reverse[r] = s[s0]
			r++
			s0++
			fallthrough
		case 3:
			reverse[r] = s[s0]
			r++
			s0++
			fallthrough
		case 2:
			reverse[r] = s[s0]
			r++
			s0++
			fallthrough
		case 1:
			reverse[r] = s[s0]
		}
		s0 = s1
	}
	// move last encoded rune, if any
	r := 0
	switch r0 {
	case 4:
		reverse[r] = s[s0]
		r++
		s0++
		fallthrough
	case 3:
		reverse[r] = s[s0]
		r++
		s0++
		fallthrough
	case 2:
		reverse[r] = s[s0]
		r++
		s0++
		fallthrough
	case 1:
		reverse[r] = s[s0]
	}
	return string(reverse)
}

// StringPrettyFast behaves as above
func StringPrettyFast(s string) string {
	s0 := 0
	s1 := 0
	r0 := len(s)
	reverse := make([]byte, r0, r0)
	for s1 = range s {
		// move the previous encoded rune, if any
		r0 -= s1 - s0
		r := r0
		for i := s0; i < s1; i++ {
			reverse[r] = s[i]
			r++
		}
		s0 = s1
	}
	// move last encoded rune, if any
	for i := 0; i < r0; i++ {
		reverse[i] = s[s0+i]
	}
	return string(reverse)
}

// StringCleanerButSlower behaves as above but uses copy, and is slower.
func StringCleanerButSlower(s string) string {
	var s0, s1 int
	r0 := len(s)
	r1 := len(s)
	reverse := make([]byte, r1, r1)
	for s1 = range s {
		if s1 > 0 { // move the previous encoded rune
			r0 -= s1 - s0
			copy(reverse[r0:r1], s[s0:s1])
			r1 = r0
			s0 = s1
		}
	}
	// move last encoded rune, if any
	copy(reverse[0:r1], s[s0:])
	return string(reverse)
}

// StringSlowerButSimpler is an alternative version that behaves differently
// in some cases. In this version, invalid utf-8 bytes are reversed as
// unicode replacement characters
func StringSlowerButSimpler(s string) string {
	r := []rune(s)
	j := len(r) - 1
	for i := 0; i < j; i++ {
		r[i], r[j] = r[j], r[i]
		j--
	}
	return string(r)
}
