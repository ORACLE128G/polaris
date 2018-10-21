package main

import (
	"flag"
	"fmt"
	"log"
	"polaris/crawler-distributed/persist"
	"polaris/crawler-distributed/rpc"
)

func serveRpc(host, index string) error {
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Index: index,
		})
}

var port = flag.Int("port", 0,
	"The port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("Must specify a port ")
		return
	}
	serveRpc(fmt.Sprintf(":%d", *port), "polaris")
}
