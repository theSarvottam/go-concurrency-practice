package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	fmt.Printf("the value of i : %v\n", i)
		// }()

		go func(i int) {
			defer wg.Done()
			fmt.Printf("the value of i : %v\n", i)
		}(i)
	}

	wg.Wait()
}
