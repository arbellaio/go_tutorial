package main

import "fmt"

func main() {

	// how slice works = if array is defined of particular size and than you want something of adjustable size than slice is used
	// slice will create new array out of initialized array with old and new data added to it

	shorthandExtendableSlice := []int{20, 30, 40} // this is also a slice
	fmt.Printf("Slice  : %v\n", shorthandExtendableSlice)
	var shorthandExtendableSliceDup = shorthandExtendableSlice
	fmt.Printf("Slice Duplicate : %v\n", shorthandExtendableSliceDup)

	//change in duplicate slice changes original slice as well
	shorthandExtendableSliceDup[0] = 60
	fmt.Printf("Slice  : %v\n", shorthandExtendableSlice)
	fmt.Printf("Slice Duplicate : %v\n", shorthandExtendableSliceDup)

	fmt.Printf("Length of slice %v\n",len(shorthandExtendableSlice))
	fmt.Printf("Capacity of slice %v\n",cap(shorthandExtendableSlice))
	//creating slice with make function

	makeSliceShorthand := make([] int, 3)

	fmt.Printf("Slice  : %v\n", makeSliceShorthand)

	fmt.Printf("Length of slice %v\n",len(makeSliceShorthand))
	fmt.Printf("Capacity of slice %v\n",cap(makeSliceShorthand))

	// creating custom capacity of slice

	makeSliceShorthandCustomCapacity := make([] int, 3, 10)

	fmt.Printf("Slice  : %v\n", makeSliceShorthandCustomCapacity)

	fmt.Printf("Length of slice %v\n",len(makeSliceShorthandCustomCapacity))
	fmt.Printf("Capacity of slice %v\n",cap(makeSliceShorthandCustomCapacity))


	//appending the slice
	sliceAppend := []int{20,40,60}
	fmt.Printf("Appendable Slice : %v\n", sliceAppend)

	appendedSlice := append(sliceAppend, 4,9)
	fmt.Printf("Appended Slice : %v\n", appendedSlice)

	removingElementAndAppending := append(sliceAppend[1:],9)
	fmt.Printf("Removed Element And Appended Slice : %v\n", removingElementAndAppending)

	//appending to slices
	sliceA := [] int {20,30,40}
	sliceB := append(sliceA, 50,60,70)
	sliceC := append(sliceB, sliceA...) //spread values of a appending B in C
	fmt.Printf("Two Appended Slice To Form Third Slice %v\n", sliceC)

}
