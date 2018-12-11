package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := 0
	for x > 0 {
		count++
		x = x & (x - 1)
	}
	return count
}
