package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func portal1(channel1 chan string)  {
	time.Sleep(2 * time.Second)
	channel1 <- "Welcome to channel1"
}

func portal2(channel2 chan string)  {
	time.Sleep(2 * time.Second)
	channel2 <- "Welcome to channel2"
}


func main() {

	//Select Statement
	R1 := make(chan string)
	R2 := make(chan string)

	go portal1(R1)
	go portal2(R2)

	select {
	case op1 := <-R1:
		fmt.Println(op1)
	case op2 := <-R2:
		fmt.Println(op2)
	}

	//ch := make(chan int, 100)
	//wg.Add(2)

	//looping through channel to get all data and case where channel is closed
	//go func(ch <-chan int) {
	//	for {
	//		if i, ok := <-ch; ok {
	//			fmt.Println(i)
	//		} else {
	//			break
	//		}
	//	}
	//	wg.Done()
	//}(ch)
	//
	//go func(ch chan<- int) {
	//	i := 12
	//	ch <- i
	//	i = 29
	//	ch <- i
	//	close(ch) // to stop data check in range for receiving data
	//	wg.Done()
	//}(ch)

	//looping through channel to get all data
	//go func(ch <- chan int) {
	//	for i := range ch {
	//		fmt.Println(i)
	//	}
	//	wg.Done()
	//}(ch)
	//
	//go func(ch chan <- int) {
	//	i:= 12
	//	ch <- i
	//	i= 29
	//	ch <- i
	//	close(ch) // to stop data check in range for receiving data
	//	wg.Done()
	//}(ch)

	// buffer for no of data in channels
	//go func(ch <- chan int) {
	//	i:= <- ch
	//	fmt.Println(i)
	//	i = <- ch
	//	fmt.Println(i)
	//	wg.Done()
	//}(ch)
	//
	//go func(ch chan <- int) {
	//	i:= 12
	//	ch <- i
	//	i= 29
	//	ch <- i
	//	wg.Done()
	//}(ch)

	// restricting go routine to send and receive only data
	//go func(ch <- chan int) {
	//	i:= <- ch
	//	fmt.Println(i)
	//	wg.Done()
	//}(ch)
	//
	//go func(ch chan <- int) {
	//	i:= 12
	//	ch <- i
	//	wg.Done()
	//}(ch)

	//example of receiving and sending the data
	//go func() {
	//	i:= <- ch
	//	fmt.Println(i)
	//	ch <- 27
	//	wg.Done()
	//}()
	//
	//go func() {
	//	i:= 12
	//	ch <- i
	//	fmt.Println(<-ch)
	//	wg.Done()
	//}()

	wg.Wait()

}
