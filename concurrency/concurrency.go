package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup is used to wait for the program to finish goroutines.
func SyncWait() {
	var wg sync.WaitGroup

	sum10 := func() {
		defer wg.Done()
		n := 0
		time.Sleep(1 * time.Second)
		for i := 0; i < 10; i++ {
			n += i
		}
		fmt.Printf("sum10 %v\n", n)
	}

	prod10 := func() {
		defer wg.Done()
		n := 1
		time.Sleep(2 * time.Second)
		for i := 1; i < 10; i++ {
			n *= i
		}
		fmt.Printf("prod10 %v\n", n)
	}

	wg.Add(2)
	go sum10()
	go prod10()
	wg.Wait()

	fmt.Println("Done!")
}
