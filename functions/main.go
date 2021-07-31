package main

import "fmt"

type rectangle struct {
	width,height int
}

func main() {

	//method example there is different between method and functions
	//functions are of specific class and can be executed by calling them where
	// as methods belong to structs objects of sorts
	rect := rectangle{
		height: 20,
		width: 20,
	}
	var areaRect = rect.area()
	fmt.Println(areaRect)

	Runable()
}

func (rect rectangle) area() int  {
	return rect.width * rect.height
}

func Runable()  {
	//Anonymous function with variable
	//fun := func() {
	//	fmt.Println("Anonymous function with variable")
	//}
	//fun()
	//*Cannot call anonymous function before declaring it

	////Anonymous function
	//func() {
	//	fmt.Println("This anonymous function")
	//}()

	//msg := "Hello User"
	//writeChangeOrderMessage(msg)

	//writeAssignPointerMessage(&msg)

	//fmt.Println(msg)

	//var sumTotal = sumWithStringWithReturn(msg, 1, 2, 3, 4, 5, 6)

	//fmt.Println(sumTotal)

	//var value, error = divideWithError(6, 0)
	//if error != nil {
	//	fmt.Println(error)
	//	return
	//}
	//fmt.Println(value)
}


//Divide with Error
func divideWithError(num1, num2 float64) (float64, error) {
	if num2 == 0 || num2 == 0.0 {
		return 0.0, fmt.Errorf("Can not divide by zero")
	}
	return num1 / num2, nil
}

//divide and collapsing types with same type
func divide(num1, num2 float64) float64 {
	return num1 / num2
}

// with return type
func sumWithStringWithReturn(msg string, values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	fmt.Println(total)
	fmt.Println(msg)
	return total
}

//variadic function with more parameters in this case add params at start
func sumWithString(msg string, values ...int) {
	total := 0
	for _, value := range values {
		total += value
	}
	//fmt.Println(total)
	fmt.Println(msg)
}

//variadic functions ... functions which can accept multiple params
//of same type

func sum(values ...int) {
	total := 0
	for _, value := range values {
		total += value
	}
	fmt.Println(total)
}

func writeAssignPointerMessageTwoParams(msg *string, msg1 string) {
	*msg = "Hello Youtube"
	fmt.Println(*msg)
}

func writeAssignPointerMessage(msg *string) {
	*msg = "Hello Youtube"
	fmt.Println(*msg)
}

func writeMessage(msg string) {
	fmt.Println(msg)
	msg = "Hello Youtube"
}

func writeChangeOrderMessage(msg string) {
	msg = "Hello Youtube"
	fmt.Println(msg)
}
