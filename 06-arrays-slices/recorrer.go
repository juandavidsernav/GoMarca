package main

import "fmt"

func main() {
	ciudades := []string{"Tokyo", "Lleida", "Paris", "Madrid"}

	for i, ciudad := range ciudades {
		fmt.Printf("[%d] %s\n", i, ciudad)
	}
}
