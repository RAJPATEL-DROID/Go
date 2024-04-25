package main

import (
	"example.com/go/practice/GoBasics"
	"fmt"
	"time"
)

// global variable
var dogsName = "rock"

func main() {
	//callChat()
	//oneToTen()
	//switchDemo()
	//arrayExample()
	//GoBasics.Datatype()
	//GoBasics.ArrayExample()

	GoBasics.FmtDemo()
}

func callChat() {
	fmt.Println("Hello chat! How are you?")
	var reply string
	fmt.Scanf("%s", &reply)
	fmt.Println("Chat : " + reply)
}

func oneToTen() {

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

	// Supported from Go 1.22
	//for j := range 11 {
	//	fmt.Println(j)
	//	j++
	//}
}

func switchDemo() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("not a number")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
