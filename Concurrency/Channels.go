package Concurrency

import "fmt"

func ChannelDemo() {
	var c1 chan int
	fmt.Println(c1) // => nil (its zero value is nil)

	c1 = make(chan int)
	fmt.Println(c1)

	//c2 := make(chan int)
	//_ = c2
	//
	//// Declaring and initilizing a RECEIVE-ONLY channel
	//c3 := make(<-chan string)
	//
	//// Declaring and initilizing a SEND-ONLY channel
	//c4 := make(chan<- string)

	//fmt.Printf("%T, %T, %T\n", c1, c3, c4) // => chan int, <-chan string, chan<- string

	c1 <- 10

	num := <-c1
	_ = num

	// Waiting for a value to be sent into the channel and print out that value
	fmt.Println(<-c1)

	close(c1)
}
