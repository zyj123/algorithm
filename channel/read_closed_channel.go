package channel

import (
	"fmt"
	"sync"
	"time"
)

func channel() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan struct{})
	go consumer(wg, ch)
	time.Sleep(3 * time.Second)
	close(ch)

	time.Sleep(3 * time.Second)
	wg.Wait()

}

func consumer(wg *sync.WaitGroup, ch chan struct{}) {
	fmt.Println(<-ch)
	fmt.Println("fin...")

	wg.Done()
}
