package main

import (
	"log"
	"net/http"
)

func main() {
	_, err := http.Get("http://127.0.0.1:8800")
	if err != nil {
		log.Printf("get error: %v\n", err)
		return
	}
}
