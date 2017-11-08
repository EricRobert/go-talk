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
	c     chan *http.Request
}

func NewModel() *Model {
	m := Model{
		c: make(chan *http.Request),
	}
	go m.process()
	return &m
}

func (m *Model) process() {
	for _ = range m.c {
		m.count++
		time.Sleep(time.Millisecond * 100)
	}
}

func (m *Model) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.c <- r
	io.WriteString(w, "hello")
}

// END OMIT

func main() {
	m := NewModel()

	go http.ListenAndServe(":8080", m)

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
