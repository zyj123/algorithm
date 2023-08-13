package p

import (
	"fmt"
	"sync"
)

func Print() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println("a")
			ch2 <- 1
		}
		//<-ch1
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch2
			fmt.Println("b")
			ch3 <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch3
			fmt.Println("c")
			ch1 <- 1
		}
	}()

	ch1 <- 1
	wg.Wait()

}
