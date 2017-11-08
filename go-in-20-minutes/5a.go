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
	n := 10

	// START OMIT
	for i := 0; i < n; i++ {
		go func() {
			r, err := http.Get("http://localhost:8080")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(i, r.StatusCode)
		}()
	}
	// END OMIT

	fmt.Println(time.Since(t))
}
