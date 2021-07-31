package main

import "fmt"

func main() {

	var amounts [3]int = [3]int{20, 30, 40}
	fmt.Printf("Amounts : %v\n", amounts)

	amountShorthand := [3]int{20, 30, 40}
	fmt.Printf("Amounts : %v\n", amountShorthand)




	//example of array pointer in GoLang
	amountItems := [3]int{10, 20, 30}
	var amountPointArray = &amountItems
	fmt.Printf("Pointed Array : %v\n", amountPointArray)

	//custom sized array
	customSizedArray := [...]int{10, 20, 30}
	fmt.Printf("Custom Sized Array : %v\n", customSizedArray)

	//copying entire array
	array1 := [...]int{10, 20, 30, 40, 50, 60, 70}
	array2 := array1[:]
	fmt.Printf("Copied Array : %v\n", array2)

	array3 := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	array4 := array3[2:]  // this will get array items after 2 index of array
	array5 := array3[:8]  // this will get array items before 8 index of array
	array6 := array3[2:8] // this will get array items after 2 index and before 8 index of array

	fmt.Printf("Copied  Array 3 : %v\n", array3)
	fmt.Printf("Copied Segmented Array 4 : %v\n", array4)
	fmt.Printf("Copied Segmented Array 5 : %v\n", array5)
	fmt.Printf("Copied Segmented Array 6 : %v\n", array6)


	var identityMatrix[3][3] int = [3][3] int {
	[3] int{1, 0, 0},
	[3] int{0, 1, 0},
	[3] int{0, 0, 1},
	}

	fmt.Printf("Multi-Dimensional Array :%v\n", identityMatrix)
	identityMatrix[0][0] = 9
	identityMatrix[1][2] = 7
	fmt.Printf("Multi-Dimensional Array :%v\n", identityMatrix)

}
