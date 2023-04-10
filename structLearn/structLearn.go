package main

import "fmt"

type PersonalDetails struct {
	Name    string
	Surname string
	age     int
}

func showDetails(p *PersonalDetails) string {
	return fmt.Sprintf("My name is %v\n My fullName is %v %v\n I am %v year old", p.Name, p.Name, p.Surname, p.age)
}

func main() {
	p := &PersonalDetails{
		Name:    "Jitendra",
		Surname: "Wadle",
		age:     35,
	}
	show := showDetails(p)
	fmt.Println(show)

}
