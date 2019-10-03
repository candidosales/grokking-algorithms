package main

import "fmt"

func main() {
	array := []int{4, 1, 8, 3, 10, 53, 2, 99, 34}
	selectionSort(array)
	fmt.Println(array)
}

func selectionSort(array []int) {
	var n = len(array)

	for i := 0; i < n; i++ {
		var smallestIndex = i
		for j := i; j < n; j++ {
			if array[j] < array[smallestIndex] {
				smallestIndex = j
			}
		}
		array[i], array[smallestIndex] = array[smallestIndex], array[i]
	}
}

// func findSmallest(array []int) int {
// 	smallest := array[0]
// 	smallest_index := 0

// 	for i := 1; i < len(array); i++ {
// 		if array[i] < smallest {
// 			smallest = array[i]
// 			smallest_index = i
// 		}
// 	}
// 	return smallest_index
// }
