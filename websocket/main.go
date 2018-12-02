package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrade = websocket.Upgrader{
		// allow cross
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wdHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, p); err != nil {
			goto ERR
		}
		fmt.Printf("got web socket message: %s \n", string(p))
	}

ERR:
	conn.Close()
}

func main() {

	http.HandleFunc("/ws", wdHandle)
	http.ListenAndServe(":8080", nil)

}
