package DataTypes

import (
	"encoding/json"
	"fmt"
)

func JsonDemo() {

	type Message struct {
		Name, Body string
		count      int64 // Won't be encoded as it's not exported type string
	}

	m := Message{"Raj", "Hi", 1223}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	//var m1 interface{}

	var m1 Message
	err = json.Unmarshal(b, &m1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m1)
}
