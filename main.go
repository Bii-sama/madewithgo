package main

import "fmt"



func main()  {

	myBill := newBill("Bidemi's Bill")

	myBill.addItem("chicken", 3.50)
	myBill.addItem("bacon", 7.50)
	myBill.addItem("coffee", 1.50)
	myBill.addItem("baked beans", 2.50)

	myBill.addTips(9)

	fmt.Println(myBill.format())
	
}


