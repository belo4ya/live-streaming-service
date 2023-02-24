package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RootResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(RootResponse{Name: "Streams", Version: "1"})
		if err != nil {
			log.Printf("err: %s\n", err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
