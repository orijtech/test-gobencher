package exhibit

// Counts the number of ones in an unsigned integer.
// Similar to POPCOUNT, or commonly known as Hamming distance.
func countOnes(in uint64) (count int) {
	if in == 0 {
		return 0
	}
	if in&(in-1) == 0 {
		return 1
	}
	// comment

	for {
		if in == 0 {
			return
		}
		if in&1 == 1 {
			count++
		}
		in >>= 1
	}
	return
}
