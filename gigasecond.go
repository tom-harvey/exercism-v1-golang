// Package gigasecond adds a billion seconds to a given time and returns it
package gigasecond

import "time"

// AddGigasecond needed the unnecessary type conversion removed
func AddGigasecond(t time.Time) time.Time {

	return t.Add(1e9 * time.Second)
}
