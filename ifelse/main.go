package main

import "fmt"

func main() {

	if true {
		fmt.Println("This is simple if statement")
	}

	if i := 2; i == 3 {
		fmt.Println("This is simple if statement with i assign to 2d")
	} else if i == 2 {
		fmt.Println("This is simple else if statement with i assign to 2d")
	} else {
		fmt.Println("This is simple else statement with i assign to 2d")
	}

	shoppingCart := make(map[string]int)
	shoppingCart = map[string]int{
		"Keyboard": 100,
		"Mouse":    200,
		"Laptop":   1000,
	}

	shoppingCart["Laptop"] = 2500
	shoppingCart["Monitor"] = 3000

	if _, ok := shoppingCart["Monitor"]; ok {
		fmt.Println("Monitor exist in shopping cart")
	}

	if _, ok := shoppingCart["Mic"]; ok {
		fmt.Println("Mic exist in shopping cart")
	}

	jk1 := 10
	jk2 := 0
	if jk1 > 0 && jk2 > 0 {
		fmt.Println("This greater than zero AND block")
	}
	if jk1 > 0 || jk2 > 0 {
		fmt.Println("This greater than zero OR block")
	}

	if jk1 == 0 && jk2 >= 0 {
		fmt.Println("This greater than zero AND block")
	}
	if jk1 <= 0 || jk2 < 0 {
		fmt.Println("This greater than zero OR block")
	}

	switch 2 {
	case 1:
		fmt.Println("This is 1")
	case 2:
		fmt.Println("This is 2")
	default:
		fmt.Println("This is default")
	}

	switch 1 {
	case 2, 4, 6, 8, 10:
		fmt.Println("This is Even")
	case 1, 3, 5, 7, 9, 11:
		fmt.Println("This is Odd")
	default:
		fmt.Println("This is default")

	}

	switch i := 2 + 2; i {
	case 2, 4, 6, 8, 10:
		fmt.Println("This is Even")
	case 1, 3, 5, 7, 9, 11:
		fmt.Println("This is Odd")
	default:
		fmt.Println("This is default")
	}

	it := 2 + 6
	switch {
	case it >= 8:
		fmt.Println("This is greater than 8")
	case it < 8:
		fmt.Println("This is less than 8")
	default:
		fmt.Println("This is default")
	}


	// to identify type of interface
	var i interface{} = 5

	switch i.(type) {
	case int:
		fmt.Println("This is int")
	case float64:
		fmt.Println("This is int")
	case string:
		fmt.Println("This is int")

	}


	switch i.(type) {
	case int:
		fmt.Println("Break Statement Test  ")
		fmt.Println("This is int")
		fmt.Println("This is int")
		break
	case float64:
		fmt.Println("This is int")
	case string:
		fmt.Println("This is int")

	}


	fmt.Println("Fallthrough")

	switch 1 {
	case 1:
		fmt.Println("This is 1")
		fallthrough
	case 2:
		fmt.Println("This is 2")
	case 3:
		fmt.Println("This is 3")
	default:
		fmt.Println("This is default")
	}
}
