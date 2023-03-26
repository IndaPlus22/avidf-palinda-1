package main

import (
	"fmt"
)

// sum the numbers in a and send the result on result

func sum(a []int, res chan<- int) {

	sum := 0
	for _, elem := range a {
		sum += elem
	}
	res <- sum
}

// concurrently sum the array

func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	return <-ch + <-ch
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
