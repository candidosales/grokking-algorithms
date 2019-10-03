package main

import "fmt"

var book map[string]float64
var voted map[string]bool

func main() {
	book = make(map[string]float64)
	book["apple"] = 0.67
	book["milk"] = 1.49
	book["avocado"] = 1.49
	fmt.Println(book)

	voted = make(map[string]bool)
	check_voter("tom")
	check_voter("mike")
	check_voter("mike")
}

func check_voter(name string) {
	if voted[name] {
		fmt.Println("kick tem out!")
	} else {
		voted[name] = true
		fmt.Println("let tem vote!")
	}
}
