package DataTypes

import "fmt"

func StructAndPtrDemo() {
	type Person struct {
		name    string
		age     int
		address string
	}
	//
	//var usr Person
	//fmt.Printf("%T", usr)

	var x int = 2
	ptr := &x
	fmt.Printf("%T\n", ptr)
	fmt.Printf("%v\n", ptr)

	// Function Type
	fmt.Printf("%T\n", f)

}

func f() {}
