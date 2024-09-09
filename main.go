package main

import (
	"fmt"
)

func updateName(x string) string { //this creates a copy of the name address in another memory space called x,
	x = "Kate" //thus when we want to update the name it changes only the reference not the original name variable
	return x   // to update the value we should return the value this is in case of group A
}

func updateMenu(y map[string]int) {
	y["ice cream"] = 210 //they don't need a return value
}
func main() {

	// Group A -> Strings, Ints, Floats, Booleans, Arrays, structs - "Non pointer values"

	name := "Kely"

	//updateName(name)
	name = updateName(name)

	fmt.Println(name)

	// Group B -> Slices, Maps, Functions - "Pointer Wrapper values"

	menu := map[string]int{
		"soup":  120,
		"salad": 150,
	}

	updateMenu(menu)
	fmt.Println(menu)

}
