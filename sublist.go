// Package sublist describes a relationship between two integer lists.
package sublist

type Relation string

// intsEqual returns true if the equal length integer slices have
// equal content.
func intsEqual(i1 *[]int, i2 []int) bool {
	for i := range *i1 {
		if (*i1)[i] != i2[i] {
			return false
		}
	}
	return true
}

// isSublist returns true if shorter is contained in longer. An empty shorter
// is considered to be contained in longer.
func isSublist(shorter *[]int, longer []int) bool {
	diff := len(longer) - len(*shorter)
	e := len(*shorter)
	for i := 0; i <= diff; i++ {
		if intsEqual(shorter, longer[i:e]) {
			return true
		}
		e++
	}
	return false
}

// Sublist returns "unequal" "equal" "sublist" or "superlist" to describe
// the relationship of i1 to i2.
func Sublist(i1, i2 []int) Relation {
	result := "unequal"
	if len(i1) == len(i2) {
		if intsEqual(&i1, i2) {
			result = "equal"
		}
	} else if len(i1) < len(i2) {
		if isSublist(&i1, i2) {
			result = "sublist"
		}
	} else { // len(i1) > len(i2)
		if isSublist(&i2, i1) {
			result = "superlist"
		}
	}
	return Relation(result)
}
