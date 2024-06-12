package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/joke", JokeHandler)
	http.HandleFunc("/madLib", MadLibHandler)
	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}

func JokeHandler(w http.ResponseWriter, r *http.Request) {
	joke := GetRandomJoke()
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(joke))
}

func MadLibHandler(w http.ResponseWriter, r *http.Request) {
	madLib := GetMadLib()
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(madLib))
}
