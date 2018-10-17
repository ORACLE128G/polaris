package main

import (
	"polaris/crawler-distributed/persist"
	"polaris/crawler-distributed/rpc"
)

func serveRpc(host, index string) error {
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Index: index,
		})
}
