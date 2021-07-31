package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop with multiple variables
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// for loop declaration outside for loop
	i, j := 0, 0
	for ; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println(i, j)

	// for loop declaration outside for loop separated checks
	k, l := 0, 0
	for k < 5 {
		fmt.Printf("K and L %v, %v\n", k, l)
		k, l = k+1, l+1
	}
	fmt.Printf("K and L %v, %v\n", k, l)

	// for loop break example
	m, n := 0, 0
	for m < 10 {
		fmt.Printf("M and N %v, %v\n", m, n)
		m, n = m+1, n+1
		if m == 6 && n == 6 {
			break
		}
	}
	fmt.Printf("M and N %v, %v\n", m, n)

	// This loop will continue to 10
	mn, no := 0, 0
	for mn < 10 {
		fmt.Printf("MN and NO %v, %v\n", mn, no)
		mn, no = mn+1, no+1
		if mn == 6 && no == 6 {
			continue
		}
	}
	fmt.Printf("MN and No %v, %v\n", mn, no)

	// This loop will continue to 6
	//mn, no := 0, 0
	//for mn < 10 {
	//
	//	if mn== 6 && no==6 {
	//		continue
	//	}
	//
	//	fmt.Printf("MN and NO %v, %v\n", mn, no)
	//	mn, no = mn+1, no+1
	//}
	//fmt.Printf("MN and No %v, %v\n", mn, no)

	//fmt.Println("Nested Loops")

	//nested loops
	//for i:=0; i<5; i++ {
	//	for j:=0; j<5; j++ {
	//			fmt.Println(i*j)
	//	}
	//
	//}

	//fmt.Println("Nested Loops Continue")

	//nested loops
	//for i := 0; i < 5; i++ {
	//	for j := 0; j < 5; j++ {
	//		if i*j == 0 {
	//			continue
	//		}
	//		fmt.Println(i * j)
	//	}
	//
	//}

	//fmt.Println("Nested Loops Break")
	//
	////nested loops
	//for i := 0; i < 5; i++ {
	//	for j := 0; j < 5; j++ {
	//		if i*j == 0 {
	//			break
	//		}
	//		fmt.Println(i * j)
	//	}
	//}

	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, v := range a {
		fmt.Printf("Index %v And Value %v\n", i, v)
	}

	//this way we dont need index
	c := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range c {
		fmt.Printf("Value %v\n", v)
	}

}
