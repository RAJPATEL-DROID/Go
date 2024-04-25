package Basics

import (
	"fmt"
	"strconv"
)

func Datatype() {
	// Declaring variables
	// var a int32 = 32132
	// var b int32 = 45452

	//Declaring multiple variables
	//const (
	//	a int = 5
	//	b     = 3
	//	c int = 123
	//)
	//
	//var d = a * b * c
	//
	//e := 12
	//
	//f := 12
	//
	//fmt.Println(d)
	//
	//fmt.Println(e + f)
	//
	// print one space between the two operands
	//fmt.Println("hi", dogsName)
	//var defaultInt int
	//var defaultFloat32 float32
	//var defaultFloat64 float64
	//var defaultBool bool
	//var defaultString string
	//fmt.Println(defaultInt, defaultFloat32, defaultFloat64, math.MaxFloat32, math.MaxFloat64, defaultBool, defaultString)

	// go is a pass by value
	firstName := "John"
	updateName(&firstName) // pass the address of the firstname
	fmt.Println(firstName)

	var a = complex(1, 2)
	fmt.Println(a)

	sum, mul := calc("10", "2")
	fmt.Println("Sum:", sum, "mul", mul)

}

func updateName(name *string) {
	*name = "raj"
}
func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}

func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}
