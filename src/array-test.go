package main

import (
	"fmt"
	"math/rand"
	"time"
)

var a []float64

func timec(fn func()) int64 {
	start := time.Now()
	fn()
	return time.Since(start).Milliseconds()
}

func unshift(number int) {
	for i := 0; i < number; i++ {
		a = append([]float64{rand.Float64()}, a...)
	}
}

func shift(number int) {
	for i := 0; i < number; i++ {
		if len(a) > 0 {
			a = a[1:]
		}
	}
}

func push(number int) {
	for i := 0; i < number; i++ {
		a = append(a, rand.Float64())
	}
}

func pop(number int) {
	for i := 0; i < number; i++ {
		if len(a) > 0 {
			a = a[:len(a)-1]
		}
	}
}

func get(idx int) func() {
	return func() {
		if idx >= 0 && idx < len(a) {
			return
		}
	}
}

// func get(idx int) func() float64 {
// 	return func() float64 {
// 		if idx >= 0 && idx < len(a) {
// 			return a[idx]
// 		}
// 		return 0.0
// 	}
// }

func pushArr(count int) func() {
	return func() {
		push(count)
	}
}

func popArr(count int) func() {
	return func() {
		pop(count)
	}
}

func unshiftArr(count int) func() {
	return func() {
		unshift(count)
	}
}

func shiftArr(count int) func() {
	return func() {
		shift(count)
	}
}

func main() {
	tests := []int{10, 100, 1000, 10000, 100000, 1000000, 10000000}
	fmt.Println("Testing get")
	for _, t := range tests {
		a = nil
		push(t)
		fmt.Println(t, timec(get(t-1)))
	}

	fmt.Println("push")
	for _, t := range tests {
		a = nil
		push(t)
		fmt.Println(t, timec(pushArr(1000)))
	}

	fmt.Println("pop")
	for _, t := range tests {
		a = nil
		push(t)
		fmt.Println(t, timec(popArr(1000)))
	}

	fmt.Println("unshift")
	for _, t := range tests {
		a = nil
		push(t)
		fmt.Println(t, timec(unshiftArr(1000)))
	}

	fmt.Println("shift")
	for _, t := range tests {
		a = nil
		push(t)
		fmt.Println(t, timec(shiftArr(1000)))
	}
}
