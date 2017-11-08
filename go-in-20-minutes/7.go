package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Model will predict stuff.
type Model struct {
	name  string
	count int
	c     chan *http.Request
}

// NewModel creates a new worker to predict stuff.
func NewModel(name string) *Model {
	m := Model{
		name: name,
		c:    make(chan *http.Request),
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

// STARTa OMIT
// Service will use models to predict stuff.
type Service struct {
	models []*Model
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(len(s.models))
	s.models[i].c <- r
	io.WriteString(w, "hello")
}

func main() {
	s := Service{
		models: []*Model{
			NewModel("a"),
			NewModel("b"),
			NewModel("c"),
			NewModel("d"),
		},
	}

	go http.ListenAndServe(":8080", &s)
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

	// STARTb OMIT
	total := 0
	for i, m := range s.models {
		total += m.count
		fmt.Println(i, m.name, "count:", m.count)
	}

	fmt.Println("count:", total)
	// END OMIT
}
