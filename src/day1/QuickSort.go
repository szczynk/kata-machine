package main

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr []int, low int, high int) {
	// base case
	if low >= high {
		return
	}

	pivotIdx := partition(arr, low, high)

	qs(arr, low, pivotIdx-1)
	qs(arr, pivotIdx+1, high)
}

func partition(arr []int, low int, high int) int {
	// nothing too fancy
	pivot := arr[high]

	// idx starts from low - 1
	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			// increment idx
			idx++

			// swap values
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	// point to pivot idx
	idx++

	// swap values
	arr[high] = arr[idx]
	arr[idx] = pivot

	return idx
}
