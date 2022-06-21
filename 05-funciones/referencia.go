package main

import "fmt"

func Incrementa(a *int) {
	*a++
}

func main() {
	a := 3
	fmt.Println(a)
	fmt.Println("-------")
	Incrementa(&a)
	fmt.Println(a)
}
