package main



import (
	"fmt"
)



// var score = 88.7

func updateName(x *string){
     *x = "Sasha"

	//  return x
}




// func updateMenu(y map[string]float64){
// 	y["bread"] = 2.99
// }


func main()  {

name := "Anikah"



// fmt.Println("memory block:",&name) //&var shows the memory location of the name variable.  memory block: 0xc000014070

//We can also store the memory block of a variable in another variable. Example:

m := &name
fmt.Println(name)
updateName(m)
// fmt.Println(m)
// fmt.Println("value at memory address is:", *m)// so in order to get the original value , just add *.
fmt.Println(name)

//Once you see an asterisk in a fuction





// 	sayHello("Hina-chan")

// fmt.Println(points)

	//maps- These are like objects in JavaScript

	menu :=  map[string]float64{
		"pasta": 6.87,
		"soup": 4.59,
		"salad": 3.12,
	}

	// updateMenu(menu)
	fmt.Println(menu)
	// fmt.Println(menu["salad"])

	// //looping through maps

	// for k, v  := range menu{
	// 	fmt.Println(k, "-", v)
	// }

	// //ints as keys

	// phoneNo := map[int]string{
	// 	1234567890: "Armin",
	// 	2345678901: "Mikasa",
	// 	3456789012: "Eren",
	// }

	// fmt.Println(phoneNo)
	// fmt.Println(phoneNo[2345678901])


	//Pass by values

	

}