package main

import (
	"fmt"
	"strconv"
)


var outScope = 12

func main() {

	// below is example of how variables are initialized
	var inScope int =  9
	var withOutType = 10
	shortForm := 7 //short form of var x int = 7 with assignment operator
	println(shortForm)
	println(inScope)
	fmt.Println(withOutType)

	//fmt is format
	// how to get value and type of variable through literals use %v to get value and %T (capital T with backslash \n)
	fmt.Printf("This is how to get value : %v, %T\n", shortForm, shortForm)

	//here var (y) is taking value (9) of within the scope and not 12 outside scope
	fmt.Printf("This is how to get value : %v, %T\n", inScope, inScope)

	// Pascal case will be used as public
	// for example variable 'I' can be used outside packages where as 'i' cannot be used outside package

	var Publicvar int = 21 // here I can be used in other packages
	fmt.Printf("This is how to get value : %v, %T\n", Publicvar, Publicvar)

	//type conversions
	var typeConversionVar float32
	// here we need to convert convert var 'i' to float32 by using float32(i)
	typeConversionVar = float32(withOutType)
	fmt.Printf("This is how to get value : %v, %T\n", typeConversionVar, typeConversionVar)


	//default value initialization

	var floatDefaultValue float32
	fmt.Printf("default value : %v, %T\n", floatDefaultValue, floatDefaultValue)

	// converting int to string
	var convertIntToString = 12
	var stringFloatValue = string(12)
	fmt.Printf("float to string : %v, %T\n", convertIntToString, convertIntToString)
	fmt.Printf("converted string ascii value : %v, %T\n", stringFloatValue, stringFloatValue)


	// converting int to string With StrConv
	var convertIntToStringWithStrConv = 12
	var stringFloatValueWithStrConv = strconv.Itoa(12)
	fmt.Printf("float to string : %v, %T\n", convertIntToStringWithStrConv, convertIntToStringWithStrConv)
	fmt.Printf("converted string ascii value : %v, %T\n", stringFloatValueWithStrConv, stringFloatValueWithStrConv)

}
