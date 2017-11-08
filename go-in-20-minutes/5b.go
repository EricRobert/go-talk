package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})

	go http.ListenAndServe(":8080", nil)

	t := time.Now()
	n := 100

	// START OMIT
	c := make(chan string)

	for i := 0; i < n; i++ {
		go func() {
			r, err := http.Get("http://localhost:8080")
			if err != nil {
				c <- err.Error()
				return
			}

			c <- r.Status
		}()
	}

	for i := 0; i < n; i++ {
		code := <-c
		log.Println(i, code)
	}
	// END OMIT

	fmt.Println(time.Since(t))
}
