package DataTypes

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func RuneDemo() {

	var1 := 'a'

	fmt.Printf("Type: %T, Value:%d \n", var1, var1)

	str := "ţară"

	//The len() built-in function returns the no. of bytes not runes or chars.
	fmt.Println(len(str)) // -> 6, number of bytes

	m := utf8.RuneCountInString(str)

	fmt.Println(m) // => 4, number of runes

	// by using indexes we get the byte at that position, not rune.
	fmt.Println("Byte (not rune) at position 1:", str[1]) // -> 163

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // -> Å£arÄ
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	// decoding a string rune by rune manually:
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:]) // it returns the rune in string in variable r

		fmt.Printf("%c", r) // -> ţară

		i += size
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	// decoding a string rune by rune automatically:
	for i, r := range str {
		fmt.Printf("%d -> %c", i, r) // => ţară
	}
}
