package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Server started my request")
	defer log.Println("Server ended my request")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Request processed successfully")
		w.Write([]byte("Request processed successfully"))
	case <-ctx.Done():
		http.Error(w, "Request cancel", http.StatusInternalServerError)
	}
}
