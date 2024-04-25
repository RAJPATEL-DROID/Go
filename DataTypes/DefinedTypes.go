package DataTypes

import "fmt"

type km float64
type mile float64

func NamedTypes() {
	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var x uint

	x = uint(s1) // can convert as the underlying type of s1(speed) is uint
	fmt.Println(x)

	var s3 speed = speed(x)
	fmt.Println(s3)

	var MumbaiTOGoa km = 65
	var DelhiTOMumbai mile

	DelhiTOMumbai = mile(MumbaiTOGoa) / 0.621
	fmt.Println(DelhiTOMumbai)

	// Aliases
	var a uint8 = 10
	var b byte // byte is an alias to uit8

	// even though they have different names,
	// byte and uit8 are the same type because they are aliases
	b = a // no error
	_ = b

	// declaring a new alias named second for uint
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint

	fmt.Printf("Minutes in an hour: %d\n", hour/60)
}
