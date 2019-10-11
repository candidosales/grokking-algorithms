package main

import "fmt"

func main() {
	result1 := substring("vista", "hish")
	result2 := substring("fish", "hish")
	fmt.Printf("Result 1: %s - Result 2: %s \n\n", result1, result2)

	result3 := subsequence("fish", "fosh")
	result4 := subsequence("fort", "fosh")
	fmt.Printf("Result 3: %d - Result 4: %d \n", result3, result4)

}

func createMatrix(rows, cols int) [][]int {
	cell := make([][]int, rows)
	for i := range cell {
		cell[i] = make([]int, cols)
	}

	return cell
}

func substring(a, b string) string {
	lcs := 0
	lastSubIndex := 0
	cell := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1

				if cell[i][j] > lcs {
					lcs = cell[i][j]
					lastSubIndex = i
				}
			} else {
				cell[i][j] = 0
			}
		}
	}

	return a[lastSubIndex-lcs : lastSubIndex]
}

func subsequence(a, b string) int {
	cell := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1
			} else {
				cell[i][j] = cell[i-1][j]

				if cell[i][j] < cell[i][j-1] {
					cell[i][j] = cell[i][j-1]
				}
			}
		}
	}

	return cell[len(a)][len(b)]
}
