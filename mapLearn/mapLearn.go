package main

import "fmt"

func main() {
	m := map[string]string{}
	m["name"] = "Jitendr"
	m["age"] = "35"
	m["surname"] = "wadle"

	fmt.Println("Full Name is: ", m["name"], m["surname"])
	fmt.Printf("%v was %v year old", m["name"], m["age"])

}
