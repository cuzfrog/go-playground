package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		sum := a + b
		a = b
		b = sum
		return sum
	}
}

func main() {

	f := fibonacci()
	for i := 0; i < 3; i++ {
		fmt.Println(f())
	}
}
