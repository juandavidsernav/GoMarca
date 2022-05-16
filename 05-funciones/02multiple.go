package main

import "fmt"

func MaxMin(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		return a, b
	}
}

func main() {

	max, min := MaxMin(10, 7)

	fmt.Println(max, min)

}