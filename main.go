package main

import (
	"github/golang-testing-class/fizzbuzz"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", fizzbuzz.FizzBuzzHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
