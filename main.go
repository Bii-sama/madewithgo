package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err:= r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill()bill  {
	reader := bufio.NewReader(os.Stdin) 

	name, _:= getInput("Create New Bill name :", reader)

	b := newBill(name)

	fmt.Println("Created the new bill -", b.name)

	return b
}

func promptOptions(b bill)  {
	reader := bufio.NewReader(os.Stdin)

	opt, _:= getInput("Choose option(a - add item, s - save a bill, t - add tip):", reader)

	switch opt{
	case "a":
       name, _ := getInput("Item Name: ", reader)
	   price, _ := getInput("Item Amount: ", reader)

	   p, err := strconv.ParseFloat(price, 64)

	   if err != nil{
		fmt.Println("The Price must be a number")
		promptOptions(b)
	   }

	   b.addItem(name, p)
		fmt.Println("Items added",name, price)
		promptOptions(b)


	case "t":
		tip, _ := getInput("Tip ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)

	   if err != nil{
		fmt.Println("The tip must be a number")
		promptOptions(b)
	   }

	   b.addTips(t)
		fmt.Println(tip)
		promptOptions(b)


	case "s":
		b.saveABill()
		fmt.Println("You saved the bill: ", b.name)
	default:
		fmt.Println("That is not a valid option")
		promptOptions(b)
	}

	// fmt.Println(opt)
}

//This is how to get and read data that is being passed in via the user input. It is used to read the 
//information the user types in. os.Stdin is used when the data is being passed from the terminal

func main()  {

	myBill := createBill()
	promptOptions(myBill)

	// fmt.Println(myBill)
	
}


