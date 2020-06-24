package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {
	a := 3
	b := 9

	fmt.Printf("TestAdd(%d, %d) = %d\n", a, b, Add(a, b))
}
