package main

import "fmt"

const firstGlobalConstString string = "Admin"

//grouped constants
const (
	 groupFirstConstant string = "FirstName"
	 groupLastConstant string = "LastName"
)

//enumeration of constants
const (
	iotaZero = iota
	iotaOne = iota
	iotaTwo = iota
)

const (
	enumValue1 = iota
	enumValue2
	enumValue3
)

// this enum starts with one
const (
	enumWithOneValue1 = iota + 1
	enumWithOneValue2
	enumWithOneValue3
	)

// this enum starts with one
const (
	enumSkips2And3WithOneValue1 = iota + 1
	_
	_
	enumSkips2And3WithOneValue4
	enumSkips2And3WithOneValue5
)


func main() {
	const firstConstInt int = 12
	fmt.Println(firstConstInt)

	const firstConstString string = "Hello"
	const firstConstFloat float32 = 3.14
	const firstConstBool bool = false


	// go does not enforce to use constant in code
	// cannot assign value to already assigned constant

	fmt.Println(iotaZero)
	fmt.Println(iotaOne)
	fmt.Println(iotaTwo)

	fmt.Println(enumValue1)
	fmt.Println(enumValue2)
	fmt.Println(enumValue3)


	fmt.Println(enumWithOneValue1)
	fmt.Println(enumWithOneValue2)
	fmt.Println(enumWithOneValue3)


	fmt.Println(enumSkips2And3WithOneValue1)
	fmt.Println(enumSkips2And3WithOneValue4)
	fmt.Println(enumSkips2And3WithOneValue5)


}
