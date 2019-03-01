// Package twofer has a package comment that summarizes what it's about.
package twofer

import (
	"strings"
)

// ShareWith has a comment documenting it.
func ShareWith(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
