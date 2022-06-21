package main

import "fmt"

type Cuboide struct {
	Ancho    float64
	Alto     float64
	Profundo float64
}

func main() {
	c := Cuboide{
		Ancho:    2,
		Alto:     3,
		Profundo: 2,
	}

	fmt.Println("La altura del cuboide es", c.Alto)
}
