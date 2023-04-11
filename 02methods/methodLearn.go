package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
	Address string
}

func (h *Human) ShowDetails() {
	fmt.Printf("My name is %v \n", h.Name)
	fmt.Printf("Full name in %v %v \n", h.Name, h.Surname)
	fmt.Printf("%v was %v year old \n", h.Name, h.Age)
	fmt.Printf("%v is leaving in %v", h.Name, h.Address)
}

func main() {
	fmt.Println("Learning methods")
	h := &Human{
		Name:    "Jitendra",
		Surname: "Wadle",
		Age:     35,
		Address: "MH",
	}
	h.ShowDetails()
}
