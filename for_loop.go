package main

import "fmt"

func main() {
	l := []int{1, 2, 3, 4}
	fmt.Println(l)
	for i := 0; i <= len(l); i++ {
		fmt.Println(i)
	}

}
