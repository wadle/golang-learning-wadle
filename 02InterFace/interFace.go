package main

import "fmt"

type Student interface {
	ShowStudentInfo()
}

type std struct {
	RollNumber int
	Name       string
	Grade      string
	Address    string
}

func (s *std) ShowStudentInfo() {
	fmt.Printf("Student Name %v \n", s.Name)
	fmt.Printf("%v Student of Grade %v\n", s.Name, s.Grade)
	fmt.Printf("%v is leaving at %v", s.Name, s.Address)
}

func main() {
	s := std{
		RollNumber: 11,
		Name:       "Jayesh Wadle",
		Grade:      "2D",
		Address:    "Wadagon Sheri Pune",
	}
	var myS Student
	myS = &s
	myS.ShowStudentInfo()

}
