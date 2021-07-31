package main

import (
	"fmt"
	"os"
)

//Defer := Defer means
//deferring execution of any particular
//statement at any particular moment
//deferred statements will be executed at last
//deferred statements will be executed in stack fashion

func main() {

	//fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)

	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)

	panic(" Problem")
	_,err := os.Create("newFile.txt")
	if err != nil {
		panic(err)
	}

	//f := createFile("defer.txt")
	//defer closeFile(f)
	//writeFile(f)
}

//// defer////
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data" )
}

func closeFile(f *os.File) {
 fmt.Println("closing")
 err:= f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %v\n",err)
		os.Exit(1)
	}
}

func createFile(fileName string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return f
}
//// defer ////