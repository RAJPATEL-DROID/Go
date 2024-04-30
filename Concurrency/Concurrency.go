package Concurrency

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters [5]int
}

func (c *Container) inc(name int) {
	//c.mu.Lock()
	//defer c.mu.Unlock()
	c.counters[name]++
}

func (c *Container) dec(name int) {

	c.counters[name]--

}

func (c *Container) add(i int) {
	//c.counters = append(c.counters, 23, 24, 5)
	c.counters[i]++
}

func ConcurrencyPractice() {
	c := Container{

		counters: [5]int{1, 2, 3, 4, 5},
	}

	var wg sync.WaitGroup

	doIncrement := func(n int) {

		for i := 0; i < 2; i++ {
			c.inc(n)
		}

		wg.Done()
	}
	doDecrement := func(n int) {
		for i := 0; i < 2; i++ {
			c.dec(n)
		}
		wg.Done()
	}

	addElement := func() {
		for i := 0; i < 3; i++ {
			c.add(i)
		}
		wg.Done()
	}
	wg.Add(5)

	go doIncrement(1)

	go doDecrement(2)

	go addElement()

	go addElement()

	go doIncrement(1)

	wg.Wait()

	fmt.Println(c.counters)
}
