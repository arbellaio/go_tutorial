package main

import "fmt"

func main() {

	shoppingCart := map[string]int{
		"Keyboard": 100,
		"Mouse":    20,
		"Laptop":   1500,
	}
	fmt.Printf("Shopping Cart %v\n", shoppingCart)

	shoppingCart2 := make(map[string]int)
	shoppingCart2 = map[string]int{
		"Keyboard": 100,
		"Mouse":    20,
		"Laptop":   1500,
	}
	fmt.Printf("Shopping Cart 2 %v\n", shoppingCart2)

	fmt.Printf("Length Shopping Cart  %v\n", len(shoppingCart))

	// get value of map element
	fmt.Printf("Price of Keyboard is %v\n", shoppingCart["Keyboard"])

	shoppingCart["Monitor"] = 1000

	fmt.Printf("Price of Monitor is %v\n", shoppingCart["Monitor"])
	fmt.Printf("Shopping Cart %v\n", shoppingCart)

	fmt.Printf("Price of Mobile is %v\n", shoppingCart["Mobile"])

	var cart = shoppingCart["Mobile"]
	fmt.Printf("Price of Mobile is %v\n", cart)

	var cartValueMobile, ok = shoppingCart["Mobile"] // this will give false against ok as mobile does not exist
	fmt.Printf("Price of Mobile is %v, %v\n", cartValueMobile, ok)

	_, ok = shoppingCart["Mobile"] // this will give false against ok as mobile does not exist
	fmt.Printf("Mobile Price %v\n", ok)

	// updating map
	shoppingCart3 := shoppingCart
	fmt.Printf("Shopping Cart %v\n", shoppingCart)
	fmt.Printf("Shopping Cart 3  %v\n", shoppingCart3)
	shoppingCart3["Keyboard"] = 2000
	fmt.Printf("Shopping Cart %v\n", shoppingCart)
	fmt.Printf("Shopping Cart 3 %v\n", shoppingCart3)

	// delete item in map
	delete(shoppingCart3, "Keyboard")
	fmt.Printf("Shopping Cart 3 %v\n", shoppingCart3)

}
