package main

import "fmt"



func main()  {

	myBill := newBill("Bidemi's Bill")

	myBill.addTips(9)

	fmt.Println(myBill.format())
	
}


