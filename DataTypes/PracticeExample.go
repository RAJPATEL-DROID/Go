package DataTypes

import (
	"fmt"
	"slices"
)

type structEx struct {
	i []int
}

func Practice() {

	//var i []structEx
	//
	//i = append(i, structEx{i: make([]int, 2, 5)})
	//
	//fmt.Printf("%T\n", i)
	//
	//fmt.Println(i)

	func() {
		i := 0
		i++
		defer fmt.Println(i)
		return
	}()

	var i int = 0
	arr := [2]int{12, 332}
	fmt.Printf("%p\n", i)
	fmt.Printf("%p\n", &i)
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", &arr)

	src := []int{12, 24, 48, 23}

	copySrc := make([]int, len(src))

	copy(copySrc, src)

	fmt.Printf("%p,%v, %p, %v\n", src, src[0], &src[0], *(&src[0]))

	fmt.Printf("%v,%v, %p\n", src, src[0], &src)

	var slice1 = slices.Delete(src, 0, 2)

	fmt.Printf("Size of Src: %v\n", src)

	slice1[1] = 52

	fmt.Printf("%v,%v, %p\n", src, src[0], &src)

	slice1 = append(slice1, 234)

	fmt.Printf("slice : %v\n", slice1)
	fmt.Printf("Src after appending Slice:  %v\n", src)

	slice1 = append(slice1, 2312)

	fmt.Printf("Slice : %v", slice1)

	fmt.Printf("Src after appending Slice:  %v\n", src)

	slice1 = append(slice1, 231212)

	fmt.Printf("slice : %v", slice1)
	fmt.Printf("Src after appending Slice: %v\n", src)

	fmt.Printf("capacity of Slice: %v\n", cap(slice1))

}
