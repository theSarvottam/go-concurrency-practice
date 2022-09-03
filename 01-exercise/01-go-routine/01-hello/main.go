package main

import (
	"fmt"
	"time"
)

func fun(str string) {

	for i := 0; i < 3; i++ {
		fmt.Println(str)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {

	//Direct call
	fun("Direct call")

	//go routine with function call

	go fun("go routine 1")

	// go routine with anonymous function call

	go func() {
		fun("GO routine 2")
	}()

	// go routine with  function value call

	fc := fun

	go fc("go routine 3")

	// Wait for go routine to end
	fmt.Println("Waiting for go routine to finish")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Done")

}
