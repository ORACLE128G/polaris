package main

import (
	"polaris/crawler-distributed/config"
	"polaris/crawler-distributed/persist"
	"polaris/crawler-distributed/rpc"
)

func serveRpc(host, index string) error {
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Index: index,
		})
}

func main() {
	serveRpc(config.ItemSaver0Host, "polaris")
}
