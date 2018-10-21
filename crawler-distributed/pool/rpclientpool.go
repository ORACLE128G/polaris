package pool

import (
	"log"
	"net/rpc"
	"polaris/crawler-distributed/rpc"
)

func CreateClientPool(hosts []string) chan *rpc.Client {

	var clients []*rpc.Client
	for _, h := range hosts{
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v",
				h, err)
		}
	}


	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
