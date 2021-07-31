package main

import "fmt"

type Student struct {
	name       string
	rollNumber int
	subjects   []string
}

func main() {

	student1 := Student{
		name:       "Faizan",
		rollNumber: 1312138,
		subjects:   []string{"Math", "Science", "English"},
	}
	fmt.Printf("Student Info : %v\n", student1)
	student1.name = "Faizan Aamer"
	fmt.Printf("Student Info : %v\n", student1)


	// example that structs are values types so only second location got updated
	student2 := student1
	student2.name = "Umar Farooq"

	fmt.Printf("Student 1 Info : %v\n", student1)
	fmt.Printf("Student 2 Info : %v\n", student2)

	// example of changing struct instance 2 and 1 with pointers


	student3 := &student1
	student3.name = "Umar Farooq"

	fmt.Printf("Student 1 Info : %v\n", student1)
	fmt.Printf("Student 3 Info : %v\n", student3)

}
