package OOP

import (
	"fmt"
	"math"
)

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func TypeAssertionAndTypeSwitchDemo() {

	var s shape = circle{radius: 2.5}

	fmt.Printf("%T\n", s) //interface dynamic type is circle

	// no direct access to interface's dynamic values
	// s.volume() -> error

	fmt.Printf("Circle Area:%v\n", s.area())

	fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

	//** TYPE SWITCHES **//
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)

	}
}
