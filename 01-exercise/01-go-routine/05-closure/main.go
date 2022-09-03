package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("THe vlue of i : %v \n", i)
		}()
		fmt.Println("Return from function")
		return
	}
	incr(&wg)
	wg.Wait()
	fmt.Println("done")

}
