package gobyex

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	fmt.Fprintf(w, "hello there %s", "world")
}

func headers(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println(k, v)
	}
}

func HttpServerFn() {
	http.HandleFunc("/getme", hello)
	http.HandleFunc("/headers", headers)

	http.HandleFunc("/hello", hellow)

	http.ListenAndServe(":8080", nil)
}

func hellow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("hello started")
	defer fmt.Println("hello ended")

	select {
	case <-time.After(time.Second * 10):
		fmt.Fprintf(w, "time ended\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
