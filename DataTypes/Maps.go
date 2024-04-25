package DataTypes

import "fmt"

func MapDemo() {
	balances := map[string]float64{
		"USD": 23.42,
		"INR": 51.23,
	}

	fmt.Printf("balances %v\n", balances)
	fmt.Printf("%T\n", balances)
}
