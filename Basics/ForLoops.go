package Basics

import "fmt"

func ForLoops() {

	// traditional for loop
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	i := 1
	//for i <= 10 {
	//	if i%2 == 0 {
	//		fmt.Println(i, "even")
	//	} else {
	//		fmt.Println(i, "odd")
	//	}
	//	i++
	//}

	// Infinite For Loop
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	// Range over Int -> Supported from Go 1.22
	//for j := range 11 {
	//	fmt.Println(j)
	//	j++
	//}
}
