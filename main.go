package main

import "fmt"

// Add returns the sum of the two supplied ints
func Add(one int, two int) int {
	return one + two
}

func main() {
	fmt.Printf("%d + %d = %d\n", 5, 4, Add(5, 4))
}
