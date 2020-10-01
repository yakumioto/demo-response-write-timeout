package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := make([]byte, 200)
		rand.Read(data)
		time.Sleep(1 * time.Second)
		n, err := writer.Write(data)
		if err != nil {
			log.Printf("[E] write error: %s\n", err)
			return
		}
		log.Printf("[I] write length: %v\n", n)
	})

	s := &http.Server{
		Addr:         ":8800",
		Handler:      mux,
		WriteTimeout: 500 * time.Millisecond,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("listening error: %v\n", err)
	}
}
