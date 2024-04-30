package Concurrency

import (
	"fmt"
	"time"
)

func BufferedChannel() {
	c1 := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Println("Func go routine before sending data into the channel", i)
			c <- i // blocks the func until main() receives the data
			fmt.Println("Func go routine after sending data into the channel", i)
		}
		close(c1)
	}(c1)

	fmt.Println("main go routine sleeps for 2 seconds")
	time.Sleep(2 * time.Second)

	go func(c1 <-chan int) {
		for v := range c1 {
			fmt.Println("Main go routine recieved value : ", v)
		}
		fmt.Print("exiting....")
	}(c1)

	time.Sleep(10 * time.Second)
}
