package main

import "fmt"

func main() {
	result := quickSort([]int{10, 4, 5, 20, 3, 2, 9, 23, 1})
	fmt.Println(result)
}

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0] // 10

		less := []int{}
		greater := []int{}

		for _, num := range array[1:] {
			if pivot > num {
				less = append(less, num) // [4,5,3,2,9,1]
			} else {
				greater = append(greater, num) // [20,23]
			}
		}

		less = append(quickSort(less), pivot)
		greater = append(greater)
		return append(less, greater...)
	}
}
