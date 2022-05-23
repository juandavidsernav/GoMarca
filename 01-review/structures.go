package main

import "fmt"

/*
Golang no cuenta con un sistema que permita realizar POO pero con metodos y structs se puede lograr
algo similar.
Golang no es orientado a objetos.
*/

type Car struct{
	Model int
	Color string
	Engine CarEngine
	Insurance Insurance
}

type CarEngine struct{
	Version int
}

type Insurance interface{

}



func main() {

	car := new(Car) //Pointer
	car2 := Car{
		Model: 2019,
		Color: "RED",
		Engine: CarEngine{
			Version: 8,
		},
	} //Reference
	// car3 := make([]Car) //array

	fmt.Printf("%v", car)
	fmt.Println()
	fmt.Printf("%v", car2)
}