package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n_0, n_1 := 0, 1
	return func() int {
		output := n_0
		n_0, n_1 = n_1, (n_0 + n_1)
		return output
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
