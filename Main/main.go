package main

import (
	"example.com/go/practice/DataTypes"
	"fmt"
)

// global variable
var dogsName = "rock"

func main() {
	//callChat()

	//Basics.ForLoops()

	//Basics.SwitchDemo()

	//Basics.Datatype()

	//Basics.ArrayExample()

	//Basics.FmtDemo()

	//Basics.ConstantDemo()

	//Basics.LabelsDemo()

	//DataTypes.RuneDemo()

	//DataTypes.StringsPack()

	//DataTypes.SliceDemo()

	DataTypes.MapDemo()

	//DataTypes.StructAndPtrDemo()

	//DataTypes.Conversion()

	//DataTypes.NamedTypes()

	//Function.FuncDemo()

	//Function.ClosureDemo()

}

func callChat() {
	fmt.Println("Hello chat! How are you?")
	var reply string
	fmt.Scanf("%s", &reply)
	fmt.Println("Chat : " + reply)
}
