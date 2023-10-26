package concurrent_print

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func test() {
	fmt.Println(solution(3, 'A', 'Z'))
	fmt.Println(solution(1, 'X', 'Z'))
	fmt.Println(solution(3, 'J', 'Q'))
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
