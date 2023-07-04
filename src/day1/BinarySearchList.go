package main

func bs_list(haystack []int, needle int) bool {
	low := 0
	// ! in Go, array indices start at 0
	high := len(haystack) - 1

	for low <= high {
		middle := (low + high) / 2
		value := haystack[middle]
		if value == needle {
			return true
		} else if value < needle {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return false

}
