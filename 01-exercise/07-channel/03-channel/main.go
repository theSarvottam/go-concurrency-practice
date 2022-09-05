package main

import "fmt"

func main() {

	//ch := make(chan int)
	// change the values here and observe the output.
	// Sent all in 1 shot as buffered channel
	ch := make(chan int, 5)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Println("Sending data to Channel : ", i)
			ch <- i
		}

	}()

	// for {
	// 	v := <-ch
	// 	// if !open {
	// 	// 	break
	// 	// }
	// 	fmt.Println(v)
	// 	time.Sleep(time.Second * 1)
	// }

	// for {
	// 	v, open := <-ch
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	for v := range ch {
		fmt.Println("Recieving data from Channel : ", v)
	}
}
