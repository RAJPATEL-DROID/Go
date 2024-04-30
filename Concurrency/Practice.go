package Concurrency

import (
	"fmt"
	"sync"
	"time"
)

func sum(a, b float64, wg *sync.WaitGroup) {

	fmt.Printf("%.2f\n", a+b)

	wg.Done()
}

func Practice() {
	//wg := sync.WaitGroup{}
	//
	//wg.Add(3)
	//
	//go sum(1.23, 12.320, &wg)
	//go sum(4.23, 5.320, &wg)
	//go sum(9.2, 6.78, &wg)
	//
	//wg.Wait()

	// Recieve simple string from Channel
	ch := make(chan string)

	go func(n string) {
		time.Sleep(2 * time.Second)
		ch <- n
	}("Gophers!")

	value := <-ch
	fmt.Println("Value received from channel:", value)
}
