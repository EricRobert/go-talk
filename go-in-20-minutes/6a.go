package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// START OMIT
type Model struct {
	count int
}

func (m *Model) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.count++
	io.WriteString(w, "hello")
}

func main() {
	m := Model{}

	go http.ListenAndServe(":8080", &m)
	// END OMIT

	t := time.Now()
	n := 100
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

	fmt.Println(time.Since(t))
	fmt.Println("count:", m.count)
}
