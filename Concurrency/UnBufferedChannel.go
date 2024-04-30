package Concurrency

import (
	"fmt"
	"time"
)

func Unbuffered() {
	var c1 = make(chan int)

	c2 := make(chan int, 3)
	_ = c2

	go func(c chan int) {
		fmt.Println("Func go routine before sending data into the channel")
		c <- 10 // blocks the func until main() receives the data
		fmt.Println("Func go routine after sending data into the channel")
	}(c1)

	fmt.Println("Main go routine sleeps for 2 sec")
	time.Sleep(2 * time.Second)

	fmt.Println("main go routine start recieving data")
	d := <-c1
	fmt.Println("main go routine recieved data from channel: ", d)

	time.Sleep(time.Second)
}
