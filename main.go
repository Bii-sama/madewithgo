package main

import "fmt"

func main()  {
	// fmt.Println("Hello, World")

	var name string = "Erwin"

	// // fmt.Println(numOne)

	// // var numTwo = "three"

	// // numsix := "seven" //this method cannot be used outside a function

	// // fmt.Println(numOne, numTwo, numsix)

	// var one int = 20
	var age = 38
	// three := 100

	// fmt.Println(one, two, three)

	// //bits and memory

	// var aNumber int8 = 25 //int8        the set of all signed  8-bit integers (-128 to 127)

	// var anotherNumber int16 = 129  // int16       the set of all signed 16-bit integers (-32768 to 32767)

	// var thirdNumber uint = 253 //uint works for only positive numbers

	// //floats- Used for decimals

	// var floatOne float32 = 98.76
	floatTwo := 35.88


	// fmt.Println(floatOne, floatTwo, aNumber, anotherNumber, thirdNumber)

    // //fmt Print
	// fmt.Print("Bidemi \n")
	// fmt.Print("Bidemi \n")
	// fmt.Print("Bidemi \n")

	//fmt Println

	// fmt.Println(name ,"was", age ,"years old when he became commander")

	//formatted strings %_ = format specifier, %v = variable, %q = variable in quotes %T = Type of variable %f = float

	fmt.Printf("%v was %v years old when he became commander \n", name, age)
	fmt.Printf("%q was %v years old when he became commander \n", name, age)
	fmt.Printf("name is of type %T and age is of type %T \n", name, age)
	fmt.Printf("%v scored %0.1f in the exam \n", name, floatTwo)

	//Sprintf() - Save formatted strings
	var str = fmt.Sprintf("%v was %v years old when he became commander", name, age)
    fmt.Println("the saved string is:",str)


}