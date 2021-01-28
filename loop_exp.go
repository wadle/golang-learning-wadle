package main

import "fmt"

func main() {
	l := []int{1, 2, 3, 4}
	fmt.Println(l)
	for i, item := range l {
		fmt.Println("Index is ", i, "and value is ", item)
	}
}
