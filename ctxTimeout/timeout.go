package ctxTimeout

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func crawler() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ch := make(chan int)
	go crawl(ch)
	select {
	case <-ctx.Done():
		fmt.Println("ctx done...")
	case <-ch:
		fmt.Println("ch read...")
	}
}

func crawl(ch chan<- int) {
	resp, err := http.Get("https://www.zhiyuanhe.com")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	ch <- 1
}
