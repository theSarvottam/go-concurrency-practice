package main

import "fmt"

func main() {

	ch := make(chan int)

	go func(a, b int) {
		sum := a + b
		ch <- sum
	}(1, 2)

	result := <-ch
	fmt.Println(result)

}
