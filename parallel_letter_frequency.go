// Package letter does these things.
package letter

// ConcurrentFrequency creates a letter frequency map from a slice of strings.
func ConcurrentFrequency(strings []string) FreqMap {
	ch := make(chan FreqMap)
	for _, s := range strings {
		go func(s string) { ch <- Frequency(s) }(s)
	}
	mergeFreq := FreqMap{}
	for range strings {
		if mergeFreq == nil {
			mergeFreq = <-ch
		} else {
			d := <-ch
			for k, v := range d {
				mergeFreq[k] += v
			}
		}
	}
	return mergeFreq
}
