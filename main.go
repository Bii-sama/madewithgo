package main

import (
	"fmt"
	"math"
	"strings"
	// "sort"
)


func sayHello(name string)  {
	fmt.Printf("Hello, %v \n", name)
}

func sayGoodBye(name string)  {
	fmt.Printf("Goodbye, %v \n", name)
}

func cycleNames(n []string, f func(string))  {
	for _, value := range n{
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}


func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)

	nameSlice := strings.Split(s, " ")

	var initials []string

	for _, v := range nameSlice {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}













func main()  {

	names := []string {"Bidemi", "Dele", "John", "Mikey", "Draken", "Angelo"}

	cycleNames(names, sayHello)
	cycleNames(names, sayGoodBye)


	// sayHello("Mikasa")
	sayGoodBye("Eren")

	a1 := circleArea(15.90)
	a2 := circleArea(10.90)

	fmt.Printf("the areas are %0.2f and %0.2f \n", a1,a2)

	fn, sn := getInitials("Eren Yaeger")
	fn2, sn2 := getInitials("Kiba")

	fmt.Println(fn, sn)
	fmt.Println(fn2, sn2)
	
// 	// fmt.Println("Hello, World")

// 	// var name string = "Erwin"

// 	// // fmt.Println(numOne)

// 	// // var numTwo = "three"

// 	// // numsix := "seven" //this method cannot be used outside a function

// 	// // fmt.Println(numOne, numTwo, numsix)

// 	// var one int = 20
// 	// var age = 38
// 	// three := 100

// 	// fmt.Println(one, two, three)

// 	// //bits and memory

// 	// var aNumber int8 = 25 //int8        the set of all signed  8-bit integers (-128 to 127)

// 	// var anotherNumber int16 = 129  // int16       the set of all signed 16-bit integers (-32768 to 32767)

// 	// var thirdNumber uint = 253 //uint works for only positive numbers

// 	// //floats- Used for decimals

// 	// var floatOne float32 = 98.76
// 	// floatTwo := 35.88


// 	// fmt.Println(floatOne, floatTwo, aNumber, anotherNumber, thirdNumber)

//     // //fmt Print
// 	// fmt.Print("Bidemi \n")
// 	// fmt.Print("Bidemi \n")
// 	// fmt.Print("Bidemi \n")

// 	//fmt Println

// 	// fmt.Println(name ,"was", age ,"years old when he became commander")

// 	//formatted strings %_ = format specifier, %v = variable, %q = variable in quotes %T = Type of variable %f = float

// 	// fmt.Printf("%v was %v years old when he became commander \n", name, age)
// 	// fmt.Printf("%q was %v years old when he became commander \n", name, age)
// 	// fmt.Printf("name is of type %T and age is of type %T \n", name, age)
// 	// fmt.Printf("%v scored %0.1f in the exam \n", name, floatTwo)

// 	//Sprintf() - Save formatted strings
// 	// var str = fmt.Sprintf("%v was %v years old when he became commander", name, age)
//     // fmt.Println("the saved string is:",str)


//     //Arrays
// 	var newarray [3]int = [3]int{60, 70, 80} 
//     //length of an array in Go lang is fixed.

// 	names := [2]string {"Bidemi", "Dele"}
// 	fmt.Println(newarray, len(newarray))
// 	fmt.Println(names, len((names)))


// 	//Slices- When a number is not specified in the box brackets for the length of an array, it is known a Slice.
// 	//Slices are easy to manipulate. The length is not fixed and we can easily add and remove elements.

// 	var newarr2 = []int {23, 78}

// 	// to append to a score
//    newarr2 = append(newarr2, 67)

//    fmt.Println(newarr2)

//    //to get ranges

//    rangeOne := newarr2[0:2] //Kinda works like javascript splice
//    rangeTwo := newarr2[0:]  //start at the index stated and gets everything till the end of the array
//    rangeThree := newarr2[:2] //start at the index stated and gets everything till the beginning of the array

//    //The ranges are slices so we can append to them

//    fmt.Println(rangeOne)
//    fmt.Println(rangeTwo)
//    fmt.Println(rangeThree)

// greetings  := "Hello, Takemitchy"
// fmt.Println(strings.Contains(greetings, "Eren"))
// fmt.Println(strings.ReplaceAll(greetings, "Takemitchy", "Eren"))
// fmt.Println(strings.ToUpper(greetings))
// fmt.Println(strings.Split(greetings, " "))//This technically converts to array

//this strings package functions do not replace the original value. It just returns a new/different one. See below
// fmt.Println("original value:", greetings)
// fmt.Println(strings.Index(greetings, "chy"))


// numbers := []int{45,98,37,26,10,89,73,62,41,50}

// sort.Ints(numbers)//The sort package/method will always alter the original slice
// fmt.Println(numbers)

// findIndex := sort.SearchInts(numbers, 50)
// fmt.Println(findIndex)

// names := []string {"Bidemi", "Dele", "John", "Mikey", "Draken", "Angelo"}

// sort.Strings(names)
// fmt.Println(names)

// fmt.Println(sort.SearchStrings(names, "Draken"))

// x := 0

// for x < 5{
// 	fmt.Println("value of x is:", x) //JavaScript while loop
// 	x++
// }

// for i:= 0; i<10; i++{
// 	fmt.Println("value of i is:", i)
// }


// sort.Strings(names)

// for i:= 0; i<len(names); i++{
// 	fmt.Println("the position of ", names[i], "is ", sort.SearchStrings(names,names[i]))
// }

// for index, value := range names{
// 	fmt.Printf("the position of %v is %v \n ", value, index)
// }

// for _, value := range names{
// 	fmt.Printf("the position of %v  \n ", value)// using the _ works when we do not want to use a particular variable
// }


//Booleans and conditionals

// age := 45

// fmt.Println(age <= 50)
// fmt.Println(age == 50)
// fmt.Println(age != 50)
// fmt.Println(age >= 50)

// if age < 30 {
// 	fmt.Println("Too young")
// }else if age > 45{
// 	fmt.Println("Too Old")
// }else{
// 	fmt.Println("Just right")
// }


// names := []string {"Bidemi", "Dele", "John", "Mikey", "Draken", "Angelo"}

// for index, value := range names {
//    if index == -1{
//     fmt.Println("continuing at position", index)
// 	continue
//    }

//    if index > 2{
// 	fmt.Println("breaking at position", index)
// 	break
//    }

//    fmt.Printf("the value at position %v is %v \n", index, value)
// }


}