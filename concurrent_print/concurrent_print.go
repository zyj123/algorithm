package concurrent_print

import (
	"fmt"
	"sync"
	"time"
)

func Test() {
	p := NewPrinter()
	for i := 0; i < 10; i++ {
		i := i
		p.Write(i)
	}

	p.Print(111)
	for i := 0; i < 10; i++ {
		i := i
		p.Write(i)
	}

	p.Stop()
}

type Printer struct {
	ch chan int
	wg *sync.WaitGroup
}

func NewPrinter() Printer {
	return Printer{
		ch: make(chan int, 100),
		wg: &sync.WaitGroup{},
	}
}

func (p *Printer) Write(i int) {
	p.ch <- i
}

func (p *Printer) Print(id int) {
	p.wg.Add(1)
	go p.print(id)
}

func (p *Printer) print(id int) {
	defer func() {
		p.wg.Done()
	}()
	for {
		time.Sleep(1 * time.Second)
		v, ok := <-p.ch
		if ok {
			fmt.Printf("printer %d print: %d\n", id, v)
		} else {
			break
		}
	}
	fmt.Println("print fin...")
}

func (p *Printer) Stop() {
	close(p.ch)
	p.wg.Wait()
	fmt.Println("stoped")
}
