package http_server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var once sync.Once
var jobQueue chan job

func init() {
	once.Do(func() {
		jobQueue = make(chan job, 3)
	})
	go do()
}

type job struct {
	w      http.ResponseWriter
	r      *http.Request
	cancel context.CancelFunc
}

func Server() {
	http.HandleFunc("/ping", ping)
	fmt.Println("listen on 8989...")
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		panic(err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	myJob := job{
		w:      w,
		r:      r,
		cancel: cancel,
	}
	jobQueue <- myJob
	<-ctx.Done()
	return
}

func do() {
	i := 0
	for {
		job := <-jobQueue
		time.Sleep(2 * time.Second)
		_, err := fmt.Fprintf(job.w, "hello "+strconv.Itoa(i))
		if err != nil {
			println("handle err:", err)
		}
		job.cancel()
		i++
	}
}
