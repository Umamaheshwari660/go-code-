package main

import "gotask/fastfood"

func main() {

	mcd := fastfood.McDonalds{City: "Hyderabad"}
	kfc := fastfood.KFC{City: "Bangalore"}
	bk := fastfood.BurgerKing{City: "Chennai"}

	fastfood.Operate(mcd)
	fastfood.Operate(kfc)
	fastfood.Operate(bk)
}


