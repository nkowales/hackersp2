// Hackers in the Bazaar Project 2
// Nathan Kowaleski, Shaquille Johnson
// 3/9/17

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello")

}
func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
