package main

import "fmt"

// go language supports composition which means it
//does not support inheritance
//It does not give you is-a relationship
//It gives you has-a relationship

type Processor struct {
	name string
	core int
}

type Memory struct {
	memoryCapacity int
	memoryType string
}

type Computer struct {
	processor Processor
	memory Memory
	price int
}

func main() {

	computer := Computer{
 	processor: Processor{
		core: 2,
		name: "Core 2 Duo",
	},
	memory: Memory{
 		memoryCapacity: 2,
 		memoryType: "DDR4",
	},
	price: 2000,
 }
 fmt.Printf("Computer Info : %v\n", computer)
}
