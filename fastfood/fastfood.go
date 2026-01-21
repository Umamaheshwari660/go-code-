package fastfood

import "fmt"

// Interface (Waiter Rules)


type Franchise interface {
	PrepareFood() string
	Menu() []string
	Location() string
}


// ---------------- McDonalds ----------------


type McDonalds struct {
	City string
}

func (m McDonalds) PrepareFood() string {
	return "Preparing Big Mac & Fries"
}
func (m McDonalds) Menu() []string {
	return []string{"Big Mac", "Fries", "McFlurry"}
}
func (m McDonalds) Location() string {
	return m.City
}


// ---------------- KFC ----------------


type KFC struct {
	City string
}

func (k KFC) PrepareFood() string {
	return "Preparing Fried Chicken"
}
func (k KFC) Menu() []string {
	return []string{"Fried Chicken", "Zinger Burger"}
}
func (k KFC) Location() string {
	return k.City
}


// ---------------- Burger King ----------------


type BurgerKing struct {
	City string
}

func (b BurgerKing) PrepareFood() string {
	return "Preparing Whopper"
}
func (b BurgerKing) Menu() []string {
	return []string{"Whopper", "Onion Rings"}
}
func (b BurgerKing) Location() string {
	return b.City
}


// Manager / Common Function


func Operate(f Franchise) {
	fmt.Println("City:", f.Location())
	fmt.Println("Menu:", f.Menu())
	fmt.Println("Status:", f.PrepareFood())
	fmt.Println("--------------------")
}
