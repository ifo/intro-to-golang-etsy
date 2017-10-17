package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/json", jsonHandler)

	fmt.Println("Server starting on port 3030")
	err := http.ListenAndServe(":3030", nil)
	log.Fatal(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This in the index page!"))
}

type JSONResponse struct {
	Name string `json:"name"`
	True bool   `json:"true"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	jr := JSONResponse{Name: "Steeeve", True: true}

	jsonBytes, err := json.Marshal(jr)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("something went wrong"))
	}

	w.Write(jsonBytes)
}
