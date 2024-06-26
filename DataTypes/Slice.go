package DataTypes

import "fmt"

func SliceDemo() {

	//var slices = []string{"Chennai", "Bengaluru"}
	//fmt.Println(slices)
	//fmt.Printf("%T\n", slices)
	//
	//var s []float32
	//printSlice(s)
	//
	//var s1 []float32
	//_ = s1

	//// append works on nil slices.
	//s = append(s, 0)
	//printSlice(s)
	//fmt.Println(cap(s))
	
	//// The slice grows as needed.
	//s = append(s, 1)
	//printSlice(s)
	//
	//// We can add more than one element at a time.
	//s = append(s, 2)
	//printSlice(s)
	//
	//s = append(s, 42, 45)
	//printSlice(s)
	//
	//s1 = append(s1, 12, 3, 23, 1)
	//printSlice(s1)
	//
	//s1 = append(s1, 23) // capacity will be doubled)
	//printSlice(s1)

	//mySlice := []int64{1.2, 5.6}
	//
	////mySlice[2] = 6
	//
	////a := 10
	//var a = 10
	//fmt.Printf("%T\n", a)
	//mySlice[0] = a
	//
	//mySlice[3] = 10.10
	//
	////mySlice = append(mySlice, a)
	//
	//fmt.Println(mySlice)

	var cities []string

	fmt.Println("cities : ", cities == nil)

	fmt.Printf("cities : %#v\n", cities)
	cities = append(cities, "Amsterland")
	fmt.Printf("cities : %#v\n", cities)

	numbers := []int{2, 3, 4, 6}
	fmt.Printf("numbers : %#v\n", numbers)

	nums := make([]int, 2)
	fmt.Println(nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)
	fmt.Printf("%T\n", friends)
	fmt.Printf("%#v\n", numbers)

	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true

	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) //false

	n1 := []int{2, 3}
	n2 := []int{4, 5}
	_, _ = n1, n2
	n1 = append(n1, n2...)
	fmt.Println(n1)

	// to compare 2 slices use a for loop to iterate over the slices and compare element by element
	// it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	var eq bool = true
	if len(a) != len(b) {
		eq = false
	}

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	// Slice headers
	// a slice expression doesn't create a new backing array. The original and the returned slice are connected!
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3] //s3, s4 share the same backing array with s1

	s3[1] = 600     // modifying the backing array so s1, s3 and s4 are in fact modified!!
	fmt.Println(s1) // -> [10 600 30 40 50]
	fmt.Println(s4) // -> [600 30]

	// when a slice is created by slicing an array, that array becomes the backing array of the new slice.
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2                       // modifying the array
	fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

	// append() function creates a complete new slice from an existing slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	// newCars doesn't share the same backing array with cars
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"                              // only cars is modified
	fmt.Println("cars:", cars, "newCars:", newCars) // => cars: [Nissan Honda Audi Range Rover] newCars: [Ford Honda
}
func printSlice(s []float32) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
