package main

import "fmt"

func main()  {
	// fmt.Println("Hello, World")

	// var numOne string = "two"

	// fmt.Println(numOne)

	// var numTwo = "three"

	// numsix := "seven" //this method cannot be used outside a function

	// fmt.Println(numOne, numTwo, numsix)

	var one int = 20
	var two = 60
	three := 100

	fmt.Println(one, two, three)

	//bits and memory

	var aNumber int8 = 25 //int8        the set of all signed  8-bit integers (-128 to 127)

	var anotherNumber int16 = 129  // int16       the set of all signed 16-bit integers (-32768 to 32767)

	var thirdNumber uint = 253 //uint works for only positive numbers

	//floats- Used for decimals

	var floatOne float32 = 98.76
	floatTwo := 35.88


	fmt.Println(floatOne, floatTwo, aNumber, anotherNumber, thirdNumber)


}