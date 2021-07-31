package main

import (
	"fmt"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}

func worker(id int, wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}


func main() {
	// running function on go routine means it will run when processor
	// feels like it
	// its running on different thread

	//message := "Hello User"
	//go func() {
	//	fmt.Println(message)
	//}()
	////go  writeMessage()
	//time.Sleep(100 * time.Millisecond)


	//Now changing value outside of go routine may cause issue
	//if go routine is taking long time to execute
	//it will change value in go routine as well eg below
	//message := "Hello User"
	//go func() {
	//	fmt.Println(message)
	//}()
	//message = "Hello Youtube"
	////go  writeMessage()
	//time.Sleep(100 * time.Millisecond)

	//Now fixing above issue by passing variable inside go routine
	//to avoid contamination

	//message := "Hello User"
	//go func(message string) {
	//	fmt.Println(message)
	//}(message)
	//message = "Hello Youtube"
	//fmt.Println(message)
	//
	////go  writeMessage()
	//time.Sleep(100 * time.Millisecond)

	// not to use time.Sleep but better approach is to use waitGroup
	//message := "Hello User"
	//wg.Add(1) //tells compiler there is only one go routine
	//go func(message string) {
	//	fmt.Println(message)
	//	wg.Done() // tells compile go routine is finished
	//}(message)
	//message = "Hello Youtube"
	//wg.Wait() // tells compiler to wait for go routine to finish

	var wg sync.WaitGroup
	for i:=1; i<=5;i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}


func writeMessage()  {
	fmt.Println("Hello")
}