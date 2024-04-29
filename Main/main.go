package main

import (
	"example.com/go/practice/Concurrency"
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

	//	DataTypes.MapDemo()

	//DataTypes.StructAndPtrDemo()

	//DataTypes.Conversion()

	//DataTypes.NamedTypes()

	//DataTypes.Practice()

	//DataTypes.JsonDemo()

	//Function.FuncDemo()

	//Function.ClosureDemo()

	//FileSystem.FileDemo()

	//FileSystem.ReadFileDemo()

	//OOP.RecieverMethods()

	//OOP.InterfaceDemo()

	//OOP.TypeAssertionAndTypeSwitchDemo()

	//OOP.GenericDemo()

	//PanicRecover.PanicDemo()

	//Concurrency.RoutineDemo()

	//Concurrency.ChannelDemo()
	Concurrency.ChannelCommunication()
}

func callChat() {
	fmt.Println("Hello chat! How are you?")
	var reply string
	fmt.Scanf("%s", &reply)
	fmt.Println("Chat : " + reply)
}
