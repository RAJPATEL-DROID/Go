package Concurrency

import (
	"fmt"
)

func square(i int, c chan int) {
	c <- i * i
	return
}

func main() {
	c := make(chan int)

	for i := 1; i < 51; i++ {
		go square(i, c)
	}

	//for i := 0; i < 50; i++ {
	//	fmt.Println(<-c)
	//}
	var cnt int = 0
	for val := range c {
		fmt.Println("recieved square from func:", val)
		cnt++
		fmt.Println(cnt)
		if cnt == 50 {
			close(c)
			break
		}
	}

}
