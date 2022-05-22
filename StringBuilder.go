package main

import (
	"fmt"
	"strings"
)

/*
String builder es una forma de concatenacion de cadenas de texto anteriormente mecionadas donde es
mas eficiente que usar unicamente el operador (+) para concatenar.
el tipo de dato se llama strings.Builder y sus funciones son Write or WriteString.
*/

func joinWithPlainString(words ...string)string{
	
	/*
	el operador difusor tambien me sirve cuando no se la cantidad de strings
	*/

	var s string

	for _, w:= range words{
		s = s+w
	}

	return s
}
/*
func join(words ...string)string{
	var sb strings.Builder
	for _,w := range words{
		i,_ := sb.WriteString(w)
		//fmt.Println(i)
	}

	return sb.String()
}*/

func join2(words ...string)string{
	var sb strings.Builder

	for _, w := range words{
		fmt.Fprintf(&sb, "%s", w)
	}
	return sb.String()
}

func main(){
	s := joinWithPlainString("Hello", " ", "My", " ","name", " ", "is Juan" )
	fmt.Println(s)

	//sb := join("Hello", " ", "My", " ","name", " ", "is Juan")
	//fmt.Println(sb)

	sb := join2("Hello", " ", "My", " ","name", " ", "is Juan")
	fmt.Println(sb)
}