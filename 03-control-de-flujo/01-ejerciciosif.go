package main

import (
	"fmt"
	"math/rand"
)

func main() {
	valor := rand.Int()

	if valor % 2 == 0{
		fmt.Println("El numero ", valor, " es par")
	}else{
		fmt.Println("El numero ", valor, " no es par")
	}
}