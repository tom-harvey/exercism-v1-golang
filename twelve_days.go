// Package twelve emits lyrics.
package twelve

const (
	pf = "On the "
	pe = "twelfth"
	pd = " day of Christmas my true love gave to me, "
	vc = "twelve Drummers Drumming, "
	vb = "eleven Pipers Piping, "
	va = "ten Lords-a-Leaping, "
	v9 = "nine Ladies Dancing, "
	v8 = "eight Maids-a-Milking, "
	v7 = "seven Swans-a-Swimming, "
	v6 = "six Geese-a-Laying, "
	v5 = "five Gold Rings, "
	v4 = "four Calling Birds, "
	v3 = "three French Hens, "
	v2 = "two Turtle Doves, and "
	v1 = "a Partridge in a Pear Tree."
	os = 0
	oe = len(pf) + os
	od = len(pe) + oe
	oc = len(pd) + od
	ob = len(vc) + oc
	oa = len(vb) + ob
	o9 = len(va) + oa
	o8 = len(v9) + o9
	o7 = len(v8) + o8
	o6 = len(v7) + o7
	o5 = len(v6) + o6
	o4 = len(v5) + o5
	o3 = len(v4) + o4
	o2 = len(v3) + o3
	o1 = len(v2) + o2
	o0 = len(v1) + o1
	s0 = pf + pe + pd + vc + vb + va + v9 + v8 + v7 + v6 +
		v5 + v4 + v3 + v2 + v1
	lastVerse = 12
)

var (
	verses = []struct {
		index            int
		saved, line, nth string
	}{
		{os, "", "", ""},
		{o1, "", v1, "first"},
		{o2, "", v2, "second"},
		{o3, "", v3, "third"},
		{o4, "", v4, "fourth"},
		{o5, "", v5, "fifth"},
		{o6, "", v6, "sixth"},
		{o7, "", v7, "seventh"},
		{o8, "", v8, "eighth"},
		{o9, "", v9, "ninth"},
		{oa, "", va, "tenth"},
		{ob, "", vb, "eleventh"},
		{oc, s0, vc, ""}, // full verse in .saved, don't need .nth
	}
	vLast     = &verses[lastVerse].saved
	savedSong = ""
)

// Song returns all verses, separated by newlines, and caches the result.
func Song() string {
	if savedSong == "" {
		for i := 1; i <= lastVerse; i++ {
			savedSong += Verse(i) + "\n"
		}
	}
	return savedSong
}

// Verse returns lyrics to one verse, without newline, and caches the result.
// It returns an empty string for bad requests.
func Verse(i int) string {
	v := &verses[0]
	if 1 <= i && i <= lastVerse {
		v = &verses[i]
		if v.saved == "" {
			b := make([]byte, 0, len(s0)) // TODO longer than needed
			b = append(b, pf...)
			b = append(b, v.nth...)
			b = append(b, pd...)
			b = append(b, (*vLast)[v.index:]...)
			v.saved = string(b)
		}
	}
	return v.saved
}
