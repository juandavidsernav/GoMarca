package main

import "fmt"

func Max(a, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

	m := Max(10, 7)
	r := 10 + Max(10, 7)

	fmt.Println(m)
	fmt.Println(r)
}
