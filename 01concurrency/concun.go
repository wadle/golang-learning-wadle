package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Learn concurrency")
	var wg sync.WaitGroup
	timeCurrent := time.Now()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go conCurrentfun(&wg)
	}

	totalTimeForConCurrent := time.Since(timeCurrent)
	fmt.Printf("Total time taken %v\n", totalTimeForConCurrent)

	normalCurrent := time.Now()
	for i := 0; i < 100; i++ {
		myFunc()
	}
	fmt.Printf("Normal Fun Total time taken %v\n", time.Since(normalCurrent))
	fmt.Printf("Concurrent Fun Total time taken %v\n", totalTimeForConCurrent)

	wg.Wait()
	// OutPut:
	//Normal Fun Total time taken 15.0317ms
	//Concurrent Fun Total time taken 505.8µs
}

func conCurrentfun(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I am running for conCurrency")
}

func myFunc() {
	fmt.Println("Normal func running")
}
