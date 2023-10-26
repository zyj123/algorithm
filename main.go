package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println(solution(3, 'A', 'Z'))
	fmt.Println(solution(1, 'X', 'Z'))
	fmt.Println(solution(3, 'J', 'Q'))
}

func solution2(interval int, start, end rune) string {
	// 请在此处编写代码
	numChan := make(chan int, 1)
	letterChan := make(chan int, 1)
	closeChan := make(chan int, 1)

	resChan := make(chan string, 1)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(interval int) {
		defer wg.Done()
		start := 1
	loop1:
		for {
			select {
			case <-numChan:
				for i := 0; i < interval; i++ {
					resChan <- strconv.Itoa(start)
					start++
				}
				letterChan <- 1
			case <-closeChan:
				break loop1
			}
		}

	}(interval)

	go func(interval int, start, end rune) {
		defer wg.Done()

	loop2:
		for {
			select {
			case <-letterChan:
				for i := 0; i < interval; i++ {
					if start > end {
						closeChan <- 1
						break loop2
					}

					resChan <- string(start)
					start++
				}
			case <-closeChan:
				break loop2
			}
		}
	}(interval, start, end)

	resArr := []string{}

loop:
	for {
		select {
		case res := <-resChan:
			resArr = append(resArr, res)
		case <-closeChan:
			break loop
		}
	}

	numChan <- 1
	wg.Wait()

	return strings.Join(resArr, "")
}

func solution(interval int, start, end rune) string {
	var sb = &strings.Builder{}

	ch1 := make(chan struct{}, interval)
	ch2 := make(chan struct{}, interval)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	num := 1
	go func() {
		defer wg.Done()
		for {
			for _ = range ch1 {

			}
			if _, ok := <-ch1; ok {
				for i := 0; i < interval; i++ {
					sb.WriteString(strconv.Itoa(num))
					num++
				}
				ch2 <- struct{}{}
			} else {
				close(ch2)
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			if _, ok := <-ch2; ok {
				for i := 0; i < interval; i++ {
					if start > end {
						close(ch1)
						return
					}
					sb.WriteByte(byte(start))
					start++
				}
				ch1 <- struct{}{}
			} else {
				break
			}
		}
	}()

	ch1 <- struct{}{}
	wg.Wait()
	return sb.String()
}
