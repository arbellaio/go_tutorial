package main

import "fmt"

func main() {
	//var boolType bool
	//fmt.Println(boolType)
	//boolType = true
	//fmt.Println(boolType)

	num1 := 20
	num2 := 10

	fmt.Printf("Num 1 + Num 2: %v\n", num1 + num2)
	fmt.Printf("Num 1 / Num 2: %v\n", num1 / num2)
	fmt.Printf("Num 1 - Num 2: %v\n", num1 - num2)
	fmt.Printf("Num 1 * Num 2: %v\n", num1 * num2)


	var complex640 complex64
	var complex641Plus2 = 1+2i

	fmt.Printf("Complex 64 0+0i : %v, %T\n", complex640, complex640)
	fmt.Printf("Complex 64 1+2i : %v, %T\n", complex641Plus2, complex641Plus2)

	var newComplexNumber = complex(2,4);
	fmt.Printf("Complex of 2 and 4 : %v, %T\n", newComplexNumber, newComplexNumber)

	//getting real and imaginary values from complex number
	fmt.Printf("Complex of 2 and 4 : %v, %T\n", real(newComplexNumber), real(newComplexNumber))
	fmt.Printf("Complex of 2 and 4 : %v, %T\n", imag(newComplexNumber), imag(newComplexNumber))


	//arithmetic of complex number
	var complexNum1 complex128 = complex(2,4)
	complexNum2 := 1+2i
	fmt.Printf("Num 1 + Num 2: %v\n", complexNum1 + complexNum2)
	fmt.Printf("Num 1 / Num 2: %v\n", complexNum1 / complexNum2)
	fmt.Printf("Num 1 - Num 2: %v\n", complexNum1 - complexNum2)
	fmt.Printf("Num 1 * Num 2: %v\n", complexNum1 * complexNum2)

	//now string type
	var string1 string = "This is 1st string "
	var string2 string = "This is 2nd string "

	fmt.Printf("%v, %T\n",string1+string2, string1+string2 )

	stringValue := 'T'
	fmt.Printf("%v, %T\n",stringValue, stringValue )

	fmt.Printf("%v, %T\n",string(stringValue), string(stringValue) )

	stringValueForArray := "This is a string"
	stringArray := []byte(stringValueForArray)
	fmt.Printf("%v, %T\n",stringArray, stringArray )
	fmt.Printf("%v, %T\n",string(stringArray), string(stringArray) )



}
