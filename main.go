package main



import (
	"fmt"
)

// var score = 88.7

func main()  {

// 	sayHello("Hina-chan")

// fmt.Println(points)

	//maps- These are like objects in JavaScript

	menu :=  map[string]float64{
		"pasta": 6.87,
		"soup": 4.59,
		"salad": 3.12,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	//looping through maps

	for k, v  := range menu{
		fmt.Println(k, "-", v)
	}

	//ints as keys

	phoneNo := map[int]string{
		1234567890: "Armin",
		2345678901: "Mikasa",
		3456789012: "Eren",
	}


}