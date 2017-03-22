package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"math/rand"
	"time"
)

type msg struct {
	Text string
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go echo(conn)
}

func echo(conn *websocket.Conn) {
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
	// for {
		// m := msg{}
		//
		// err := conn.ReadJSON(&m)
		// if err != nil {
		// 	fmt.Println("Error reading json.", err)
		// }

		fmt.Printf("Got message: %#v\n", s)

		if err := conn.WriteJSON(msg{s}); err != nil {
			fmt.Println(err)
		}
	// }
}
