package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// START OMIT
// ...

func main() {
	t := time.Now()
	n := 10

	for i := 0; i < n; i++ {
		r, err := http.Get("https://element.ai")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(i, r.StatusCode)
	}

	fmt.Println(time.Since(t))
}

// END OMIT
