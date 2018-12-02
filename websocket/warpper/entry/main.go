package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"polaris/websocket/warpper"
	"time"
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
	var wrapper *connection.Connection
	var data []byte
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	if wrapper, err = connection.InitConnetion(conn); err != nil {
		goto ERR
	}

	go func() {
		for {
			if err := wrapper.WriteMsg([]byte("heartbeats checking")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		if data, err = wrapper.ReadMsg(); err != nil {
			goto ERR
		}
		if err := wrapper.WriteMsg(data); err != nil {
			goto ERR
		}
	}

ERR:
	wrapper.Close()
}

func main() {

	http.HandleFunc("/ws", wdHandle)
	http.ListenAndServe(":8080", nil)

}
