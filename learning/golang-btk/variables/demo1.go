package variables

import "fmt"

func Demo1() {

	// fmt.Println("Hello from go!")

	//! Variables
	// var word string = "Hello from go!"
	// fmt.Println(word)
	// fmt.Println(word)
	// fmt.Println(word)
	// fmt.Println(word)
	// fmt.Println(word)

	//? Integer
	// var tax int = 15
	// fmt.Println(tax)
	// fmt.Println(100 + 100 * tax / 100)

	//? Float
	// var tax2 float32 = 0.1
	// fmt.Println(tax2)
	// fmt.Println(100 + 100 * tax2)

	// var tax3 int = 111
	// tax4 := "Ken Den"

	// fmt.Printf("Data type : %T\n", tax3)
	// fmt.Printf("Data type : %T", tax4)

	//! Boolean
	var state bool = false
	// state1 := true

	string1 := "Ken"
	string2 := "Den"

	state = string1 == string2
	fmt.Println(state)

}