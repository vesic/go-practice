package concurrency

import (
	"fmt"
	"math/rand"
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

// Communication with channels.
func Channels1() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		fmt.Println("send to c1")
		c1 <- rand.Intn(100)
	}()

	go func() {
		fmt.Println("send to c2")
		c2 <- rand.Intn(100)
	}()

	fmt.Println("rcv from c1", <-c1)
	fmt.Println("rcv from c2", <-c2)
}
