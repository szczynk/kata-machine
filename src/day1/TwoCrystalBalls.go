package main

import "math"

func two_crystal_balls(breaks []bool) int {
	jumpAmount := math.Floor(math.Sqrt(float64(len(breaks))))

	i := int(jumpAmount)
	for i < len(breaks) {
		if breaks[i] {
			break
		}
		i += int(jumpAmount)
	}

	i -= int(jumpAmount)

	j := 0
	for j < int(jumpAmount) && i < len(breaks) {
		if breaks[i] {
			return i
		}
		j++
		i++
	}

	return -1
}
