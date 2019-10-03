package main

import "fmt"

func main() {
	result := sum([]int{1, 2, 3, 4, 5})
	fmt.Println(result)

	result = sum([]int{})
	fmt.Println(result)

	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array[1:])

	result = sumRecursive(array)
	fmt.Println(result)

	result = listSize(array)
	fmt.Println(result)

	result = maxNumber([]int{6, 1, 2, 3, 4, 5})
	fmt.Println(result)
}

func sum(array []int) int {
	total := 0

	for _, num := range array {
		total += num
	}

	return total
}

func sumRecursive(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return array[0] + sumRecursive(array[1:])
}

// 1 + [2 3 4 5]
// 2 + [3 4 5]
// 3 + [4 5]
// 4 + [5]
// 5

func listSize(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return 1 + listSize(array[1:])
}

// 1 + [2 3 4 5]
// 1 + [3 4 5]
// 1 + [4 5]
// 1 + [5]
// 1

func maxNumber(array []int) int {
	if len(array) == 2 {
		if array[0] > array[1] {
			return array[0]
		}
		return array[1]
	}

	max := maxNumber(array[1:])
	if array[0] > max {
		return array[0]
	}
	return max
}

//
