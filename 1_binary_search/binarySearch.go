package main

import "fmt"

func main() {
	result, steps := binarySearch([]int{1, 2, 3, 4, 55, 54, 77, 79, 94, 95, 100}, 3)
	fmt.Printf("result=%d steps=%d \n", result, steps)

	result, steps = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}, 19)
	fmt.Printf("result=%d steps=%d \n", result, steps)

	result, steps = binarySearch([]int{1, 3, 5, 7, 9}, 3)
	fmt.Printf("result=%d steps=%d \n", result, steps)

	result, steps = binarySearch([]int{1, 3, 5, 7, 9}, -1)
	fmt.Printf("result=%d steps=%d \n", result, steps)

	var slice []int
	count := 0
	max := 128
	for count < max {
		slice = append(slice, count)
		count = count + 1
	}

	result, steps = binarySearch(slice, 3)
	fmt.Printf("result=%d steps=%d \n", result, steps)
}

func binarySearch(list []int, item int) (int, int) {

	low := 0
	high := len(list) - 1 // 4

	steps := 0

	for low <= high {

		steps = steps + 1

		mid := (low + high) / 2 // (0+4)/2 // (0+1)/2 = 0
		guess := list[mid]      // 5       // 1

		if guess == item {
			return mid, steps
		} else if guess > item {
			high = mid - 1 // 1
		} else {
			low = mid + 1 // 1
		}
	}
	return 0, steps
}

// 1.1 128 / log b (x) = y / log 2 (128) = x / 2 ^ x = 128 / 2 ^ 7 = 128
// 1.2 256 / 8
