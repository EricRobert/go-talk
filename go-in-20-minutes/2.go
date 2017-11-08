package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	t := time.Now()

	r, err := http.Get("https://element.ai")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.StatusCode, time.Since(t))
}
