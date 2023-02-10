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

// Receive values using range expression.
func Channels2() {
	var wg sync.WaitGroup
	c1 := make(chan int)
	c2 := make(chan int)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			c1 <- rand.Intn(100)
		}
		close(c1)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			c2 <- rand.Intn(100)
		}
		close(c2)
	}()

	even := func(val int) {
		fmt.Printf("even(%v) = %v\n", val, val)
	}

	odd := func(val int) {
		fmt.Printf("odd(%v) = %v\n", val, val)
	}

	for val := range c1 {
		if val%2 == 0 {
			even(val)
		} else {
			odd(val)
		}
	}

	for val := range c2 {
		if val%2 == 0 {
			even(val)
		} else {
			odd(val)
		}

	}

	wg.Wait()
}

// Execute synchronous functions asynchronously in a wrapper,
func ExecuteSyncFuncAsync() {
	chan1 := make(chan string)

	hello := func(s string) string {
		return fmt.Sprintf("Hello %s!", s)
	}

	wrapper := func(s string, c chan string) {
		c <- hello(s)
	}

	go wrapper("world", chan1)

	go func(s string, c chan string) {
		c <- hello(s)
	}("Go", chan1)

	for i := 0; i < 2; i++ {
		fmt.Println(<-chan1)
	}
}
