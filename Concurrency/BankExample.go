package Concurrency

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {

	mu.Lock()
	*b += n
	defer mu.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	*b -= n
	defer mu.Unlock()
	wg.Done()
}

func BankDemo() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &mu)
		go withdraw(&balance, i, &wg, &mu)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
