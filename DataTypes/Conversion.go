package DataTypes

import (
	"fmt"
	"strconv"
)

func Conversion() {
	s := string(99)
	fmt.Println(s)

	//s1 := string(99.123) -> throws error
	//var myStr = fmt.Sprintf("%f", 44.2)
	//fmt.Printf(myStr) // String representation of 44.2=

	var myStr = fmt.Sprintf("%d", 3232)
	fmt.Printf(myStr)

	fmt.Println(string(3232)) // give unicode character of value given

	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Printf("%T\n", f1)
	fmt.Println(f1 * 3.4)

	i4, err := strconv.Atoi("2446")
	str2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i valus is %v\n", i4, i4)
	fmt.Printf("str2 type is %T, str2 valus is %v\n", str2, str2)

}
