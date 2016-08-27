package thelm

// return the value between low and high inclusive
func minmax(low int, value int, high int) int {
	if value < low {
		return low
	} else if value > high {
		return high
	}
	return value
}
