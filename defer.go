package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hi for main")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("I`m still alive, recover")
			context.Background()
		}
	}()

	g(99)
}

func g(i int) {
	defer fmt.Println("not defer yet")
	fmt.Println(i)
	panic("going down dude")
}

func deferTest() {
	defer timeTracker(time.Now(), "main")
	fmt.Println("first")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Second")
}

func timeTracker(t time.Time, name string) {
	elapse := time.Since(t)

	fmt.Println(fmt.Sprintf("the function %s took %s", name, elapse.String()))
}
