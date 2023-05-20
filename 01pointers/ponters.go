package main

import "fmt"

func main() {
	var number = 10
	fmt.Println("Learn Pointers")
	var p *int
	fmt.Printf("Initial Value is %v\n", p)
	p = &number
	fmt.Printf("P can hold address :: %v \n", p)
	fmt.Printf("Value of P is %v", *p)
}
