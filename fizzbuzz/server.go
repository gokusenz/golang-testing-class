package fizzbuzz

import (
	"fmt"
	"net/http"
)

// FizzBuzzHandler : Fizzbuzz
func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("s")
	fmt.Fprintf(w, "%v", Fizzbuzz(param1))
}
