package main

import "fmt"

func main() {
	var hora, minuto, segundo int

	fmt.Println("HH:MM:SS?")
	fmt.Scanf("%d, %d, %d", &hora, &minuto, &segundo)
	fmt.Printf("%d horas, %d minutos, %d segundos", hora, minuto, segundo)
}
