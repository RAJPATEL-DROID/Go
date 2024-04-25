package Basics

import (
	"fmt"
)

func LabelsDemo() {

	// declaring a variable
	// there is no conflict name between variable and label
	outer := 19
	_ = outer

	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Marry"}

outer: //label, it doesn't conflict with other names
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("FOUND A FRIEND: %q at index %d\n", friend, index)
				break outer //breaking outside the outer loop which terminates
			}
		}
	}

	fmt.Println("Next instruction after the break.")

	i := 0
loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//  goto label
	// ERROR it's not permitted to jump over the declaration of x
	//  x := 5
	// label:
	//  	fmt.Println("something here")
}
