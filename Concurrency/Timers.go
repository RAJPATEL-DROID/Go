package Concurrency

import (
	"fmt"
	"time"
)

func TimerDemo() {

	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println("hello", <-timer1.C)

	//fmt.Println("Timer 1 fired")

	fmt.Printf("%T\n", timer1.C)

	timer2 := time.NewTimer(time.Second)

	go func() {

		<-timer2.C

		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()

	if stop2 {

		fmt.Println("Timer 2 stopped")

	}

	time.Sleep(2 * time.Second)
}
