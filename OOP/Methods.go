package OOP

import (
	"fmt"
	"time"
)

type names []int

func (n names) print() {
	for i, val := range n {
		fmt.Println(i, " ", val)
	}
}

type car struct {
	brand string
	price int
}

func changeCarDetails(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCarDetails1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCarDetails2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

func ReceiverMethods() {

	// Hour() isthe reciever methods for Time data type
	const day = 24 * time.Second

	fmt.Printf("%T", day)

	var hour = day.Hours()

	fmt.Printf("%T\n", hour)

	friends := names{1, 2, 3}

	friends.print()

	// names.print(friends)

	myCar := car{"Audi", 40000}
	changeCarDetails(myCar, "Renault", 20000)
	fmt.Println(myCar) // No change appears in car details

	myCar.changeCarDetails1("Kia", 100000)
	fmt.Println(myCar) // No change as passed by value

	(&myCar).changeCarDetails2("reanult", 203203)
	myCar.changeCarDetails1("reanult", 203203) // same as above( implicit & will be added)

	fmt.Println(myCar)

	var yourcar *car
	yourcar = &myCar
	fmt.Println(*yourcar)

	yourcar.changeCarDetails2("vw", 400000)
	fmt.Println(*yourcar)

	(*yourcar).changeCarDetails2("another vw", 4000023)
	fmt.Println(*yourcar)

	fmt.Println(myCar)

}
