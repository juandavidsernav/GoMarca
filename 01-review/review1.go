/*
Estructuras, punteros, manejo de errores, funciones, metodos y el tipo nil
*/

package main 

import "fmt"

func seeThisCondicional(i int)bool{
	var b bool

	if i%2 == 0{
		b = true
	}else{
		b = false
	}

	return b
}

func main(){

	a := seeThisCondicional(2)
	b := seeThisCondicional(11)
	fmt.Println(a)
	fmt.Println(b)
}