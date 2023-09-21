package http_server

import (
	"fmt"
	"net/http"
	"time"
)

func server() {
	http.HandleFunc("/", hello)

	fmt.Println("listen on 8989...")
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", w.Header())

	time.Sleep(1 * time.Second)
	n, err := fmt.Fprintf(w, "hello")
	fmt.Println(n, "  ", err)
	fmt.Println("handle finish...")
	return
}
