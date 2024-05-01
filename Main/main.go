package main

import (
	"example.com/go/practice/Concurrency"
	"example.com/go/practice/OOP"
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

	OOP.ReceiverMethods()

	//OOP.InterfaceDemo()

	//OOP.TypeAssertionAndTypeSwitchDemo()

	//OOP.GenericDemo()

	//PanicRecover.PanicDemo()

	//Concurrency.ConcurrencyPractice()

	//Concurrency.MapModification()

	//Concurrency.RoutineDemo()

	//Concurrency.ChannelDemo()

	//Concurrency.ChannelCommunication()

	//Concurrency.Unbuffered()

	//Concurrency.BufferedChannel()

	//Concurrency.Practice()

	//Concurrency.BankDemo()

	//Concurrency.SelectDemo()

	Concurrency.TimerDemo()

}

func callChat() {
	fmt.Println("Hello chat! How are you?")
	var reply string
	fmt.Scanf("%s", &reply)
	fmt.Println("Chat : " + reply)
}
