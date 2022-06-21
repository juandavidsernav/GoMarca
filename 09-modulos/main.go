package main

import (
	"fmt"

	"github.com/mariomac/analizador"
	"juan.info/go/ejemplo/hola"
)

func main() {
	fmt.Print("Como te llamas?: ")
	var nombre string
	fmt.Scanln(&nombre)
	fmt.Println(hola.ConNombre(nombre))

	analizador.PrintEstadistica(nombre)
}
