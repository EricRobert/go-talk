package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r, err := http.Get("https://element.ai")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.StatusCode)
}
