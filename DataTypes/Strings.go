package DataTypes

import (
	"fmt"
	s "strings"
)

func StringsPack() {

	// declaring a variable of type func to call the Println function easier.

	fmt.Println("Contains:  ", s.Contains("test", "es"))
	fmt.Println("Count:     ", s.Count("test", "t"))
	fmt.Println("HasPrefix: ", s.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", s.HasSuffix("test", "st"))
	fmt.Println("Index:     ", s.Index("test", "e"))
	fmt.Println("Join:      ", s.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", s.Repeat("a", 5))
	fmt.Println("Replace:   ", s.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", s.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", s.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", s.ToLower("TEST"))
	fmt.Println("ToUpper:   ", s.ToUpper("test"))

}
