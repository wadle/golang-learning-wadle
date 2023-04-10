package main

import "fmt"

type LearnInterface interface {
	ShowName()
	ShowFullName()
	ShowAge()
}

type PersonalDetails struct {
	Name    string
	Surname string
	Age     int
}

func (p *PersonalDetails) ShowName() {
	fmt.Println("My name is ", p.Name)
}

func (p *PersonalDetails) ShowFullName() {
	fmt.Println("My FullName is ", p.Name, p.Surname)
}

func (p *PersonalDetails) ShowAge() {
	fmt.Printf("%v was %v old", p.Name, p.Age)
}

func Print(in LearnInterface) {
	in.ShowName()
	in.ShowFullName()
	in.ShowAge()
}

func main() {
	p := &PersonalDetails{
		Name:    "Jitendra",
		Surname: "Wadle",
		Age:     35,
	}

	Print(p)
}
