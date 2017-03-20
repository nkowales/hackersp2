// Hackers in the Bazaar Project 2
// Nathan Kowaleski, Shaquille Johnson
// 3/19/17

package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello")

}
func main() {

	http.HandleFunc("/fortune", fortune)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func fortune(w http.ResponseWriter, r *http.Request) {
	
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	n := r1.Intn(20)
		
	var f [20]string
	f[0] = "It is certain"
	f[1] = "It is decidedly so"
	f[2] = "Without a doubt"
	f[3] = "Yes definitely"
	f[4] = "You may rely on it"
	f[5] = "As I see it, yes"
	f[6] = "Most likely"
	f[7] = "Outlook good"
	f[8] = "Yes"
	f[9] = "Signs point to yes"
	f[10] = "Reply hazy try again"
	f[11] = "Ask again later"
	f[12] = "Better not tell you now"
	f[13] = "Cannot predict now"
	f[14] = "Concentrate and ask again"
	f[15] = "Don't count on it"
	f[16] = "My reply is no"
	f[17] = "My sources say no"
	f[18] = "Outlook not so good"
	f[19] = "Very doubtful"
	
	s := f[n]
	fmt.Fprintf(w, "%s", s)
	
}
