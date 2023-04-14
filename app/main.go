package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		w.Write([]byte("Hello"))
	})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
