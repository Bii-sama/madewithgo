package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}


//make a bill


func newBill(name string)bill{
    b:= bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b

}


//Receiver functions: These are functions tha work with a struct

func (b bill)format() string  {
	
	fs := "Bill breakdown: \n"

	var total float64 = 0

	for k, v:= range b.items{
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total+=v

	}

	fs += fmt.Sprintf("%-25v....%0.2f", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v....%0.2f", "total:", total)

	return fs

}


//update bill items and tips

func (b *bill) addTips(tip float64) {
	(*b).tip = tip
	
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

