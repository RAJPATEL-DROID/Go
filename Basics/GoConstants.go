package Basics

import (
	"fmt"
	"math"
)

func ConstantDemo() {
	const days = 7

	var i int
	fmt.Println(i)
	const pi float32 = 3.14
	const secondInHour = 3600

	duration := 234
	fmt.Println("Duration in second : ", duration*secondInHour)

	// As this is variables
	//x, y := 5, 0
	// following will be caught at runtime
	//fmt.Println(x / y)

	//const x = 5
	//const y = 0
	// Following will be caught at compile time
	//fmt.Println(x / y)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min2 = -500
		// min3 and min5 will get the value and type from above constant
		min3
		min5
	)

	fmt.Println(min2 + min3 + min5)

	// Constant Rules
	// 1. Can't initiate the constant at runtime
	//const mathPow = math.Pow(2, 4)

	// 2. can't use a variable to initialise a constant
	//t := 5
	//const tc = t // this will give error,as variable are runtime values

	//const len1 = len("hello") // as a len is builtin function available at compile time
	//fmt.Println(len1)

	//const a float32 = 5.1 // typed constant
	//const b = 6.7         // untyped constant
	//
	//const c float32 = a * b
	//
	//const str = "Hello" + "good"
	//
	//const d = 5 > 10
	//fmt.Println(d)

	// In typed Constant
	//const x int = 5
	//const y float32 = 2.2 * x // not possible

	// In Untyped constant
	const x = 5
	const y = 2.2 * 5
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	var i1 int = x
	var j1 float32 = x // -> (float32(x))
	var p byte = x     // -> byte(x)
	fmt.Println(i1, j1, p)

	fmt.Printf("$%T\n", p)

	// IOTA -> predeclared identifier
	// represents successive untyped integer constants
	// value is the index of the respective ConstSpec, starting with zero
	const (
		c0 = iota
		c1
		c2
		c3 = 4
		c4
		c5 = 5
	)

	fmt.Println(c0, c1, c2, c3, c4, c5)

	const (
		a = 1 << iota
		b
		c
		d
	)

	fmt.Println(a, b, c, d)

	//const (
	//	x = -(iota + 2)
	//	_
	//	y
	//	z
	//)

	// OVERFLOW
	var x1 uint8 = 255
	x1++ // Can't catch at compile time, so it will just overflow
	fmt.Println("Overflowed value", x1)

	//a1 := int8(255 + 1) // will be caught at compile time n gives error
	//_ = a1

	var b1 int8 = 127
	fmt.Println(b1 + 1) // overflows

	b1 = -128
	b1--
	fmt.Println(b1) // overflow

	f := float32(math.MaxFloat32)
	f = f * 1.2
	fmt.Println(f) // gives +Inf as a result of overflowing

}
