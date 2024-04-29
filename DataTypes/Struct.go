package DataTypes

import (
	"fmt"
	"strings"
)

func StructAndPtrDemo() {
	type Person struct {
		name    string
		age     int
		address string
	}
	//
	//var usr Person
	//fmt.Printf("%T", usr)

	var x int = 2
	ptr := &x
	fmt.Printf("%T\n", ptr)
	fmt.Printf("%v\n", ptr)

	type book struct {
		title  string //the fields of the book struct
		author string //each field must be unique inside a struct
		year   int
	}

	type book1 struct {
		title, author string
		year, pages   int
	}

	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
	_ = bestBook

	aBook := book{title: "Just a random book"}
	fmt.Printf("%#v\n", aBook) // => main.book{title:"Just a random book", author:"", year:0}

	// pages := lastBook.pages // error -> lastBook.pages undefined (type book has no field or method pages)

	// updating a field
	bestBook.author = "The Best"
	bestBook.year = 2020
	// + with %v  prints both the field names and their values
	fmt.Printf("lastBook: %+v\n", bestBook)

	randomBook := book{title: "Random Title", author: "John Doe", year: 100}

	fmt.Println(randomBook == bestBook) // => false

	//a copy of a struct
	myBook := randomBook
	myBook.year = 2020              // modifying only myBook
	fmt.Println(myBook, randomBook) // => {Random Title John Doe 2020} {Random Title John Doe 100}

	// anonymousStruct()
}

func anonymousStruct() {
	// an anonymous struct is a struct with no explicitly
	// defined struct type alias.
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	// =>struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Muller", age:30
	fmt.Printf("%#v\n", diana)

	//** ANONYMOUS FIELDS **//

	// fields type becomes fields name.
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1) // => main.Book{string:"1984 by George Orwell", float64:10.2, bool:false}

	// mixing anonymous with named fields:
	type Employee1 struct {
		name   string
		salary int
		bool
	}

	e := Employee1{"John", 40000, false}
	fmt.Printf("%#v\n", e) // => main.Employee1{name:"John", salary:40000, bool:false}
	e.bool = true          // changing a field

	fmt.Println(strings.Repeat("#", 10))

	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 3000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   042324234,
		},
	}

	fmt.Printf("%+v\n", john)

	fmt.Printf("Employee's salary: %d\n", john.salary)

	fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // => Employee's email:jkeller@company.com

	john.contactInfo.email = "new_email@thecompany.com"

	fmt.Printf("Employee's new email address:%s\n", john.contactInfo.email)
}
