// Hackers in the Bazaar Project 2
// Nathan Kowaleski, Shaquille Johnson
// 3/19/17

package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// wsHandler implements the Handler Interface
type wsHandler struct{}

type msg struct {
	Text string
}

var upgrader = websocket.Upgrader{} // use default options

func fortune() []byte {

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
	return []byte(s)

}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, fortune())
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

// 	//Upgrading HTTP Connection to websocket connection
// 	wsConn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 	    log.Printf("error upgrading %s", err)
// 	    return
// 	}
//
// 	//handle your websockets with wsConn
// 	wsConn.WriteJSON(msg{fortune()})
// }

func main() {

	http.HandleFunc("/fortune", echo)
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":8080", nil)

}
