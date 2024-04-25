package GoBasics

import "fmt"

func ArrayExample() {
	var a [5]int8
	fmt.Println("emp: ", a)

	a[1] = 10
	fmt.Println(a[1])

	fmt.Println(len(a))

	b := [5]int8{10, 20, 30, 40, 50}
	fmt.Println(b)

	b = [...]int8{10, 20, 10, 30, 50}
	fmt.Println(b)

	// Below line will give error : no new variables on left side of :=
	//b := [...]int8{10, 20, 30, 40, 50}

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	tooD := [2][3]int{
		{1, 2, 0},
		{4, 2, 1},
	}
	tooD[0][0] = 31
	fmt.Println(tooD)

	tooD = [2][3]int{
		{1, 2, 0},
		{4, 2, 1},
	}
	tooD[1][2] = 113
	fmt.Println(tooD)

}
