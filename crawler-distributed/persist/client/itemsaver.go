package client

import (
	"log"
	"polaris/crawler-distributed/config"
	"polaris/crawler-distributed/rpc"
	"polaris/crawler/engine"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			result := ""
			err = client.Call(config.ItemSaverServiceSave, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item: %v, error: %v", item, err)
			}
		}
	}()
	return out, nil

}
