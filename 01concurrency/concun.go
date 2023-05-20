package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learn concurrency")
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go conCurrentfun(&wg)
	}
	wg.Wait()

}

func conCurrentfun(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I am running for conCurrency")
}
