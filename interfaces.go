package main

import "fmt"

/*Las interfaces en Golang son implicitas (no se tiene una palabra clave),
struct -> interface
*/

func main() {
	fmt.Println("I want a Coffee")

	var (
		i ItalianCoffeeMachine
		c ColombianCoffeeMachine
	)

	italianCoffee := GetCoffee(&i, 10)
	italianCoffee.PrintCoffee()

	colombiaCoffee := GetCoffee(&c, 25)
	colombiaCoffee.PrintCoffee()

}

type Coffee struct {
	Intensity int
	Region    string
}

func (c *Coffee) PrintCoffee() {
	fmt.Println(fmt.Sprintf("This coffee es from %s and intensity is %d", c.Region, c.Intensity))
}

type CoffeeMaker interface {
	MakeCoffee(intensty int) Coffee
}

type ItalianCoffeeMachine struct {
}

type ColombianCoffeeMachine struct {
}

func (i *ItalianCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Italy"}
}

func (c *ColombianCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Colombia"}
}

func GetCoffee(coffeeMaker CoffeeMaker, i int) Coffee {
	return coffeeMaker.MakeCoffee(i)
}
