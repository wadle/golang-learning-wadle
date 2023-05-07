package main

import "fmt"

type School struct {
	Roots    []string
	Division []string
}

func (s *School) ShowData() {
	for rootNumber, root := range s.Roots {
		fmt.Printf("RootNumber %v is %v\n", rootNumber, root)
	}

	for divNumber, div := range s.Division {
		fmt.Printf("DivisionNumber %v is %v\n", divNumber, div)
	}

}

func main() {
	fmt.Println("Methods demo")
	s := &School{
		Roots:    []string{"Root1Test1", "Root2Test2", "Root3Test3", "Root4Test4"},
		Division: []string{"Division1Test1", "Division2Test2", "Division3Test3", "Division4Test4"},
	}
	s.ShowData()

}
