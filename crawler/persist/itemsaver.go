package persist

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"github.com/olivere/elastic/config"
	"log"
	cfg "polaris/crawler-distributed/config"
	"polaris/crawler/engine"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(item, "polaris")
			if err != nil {
				log.Printf("Item Saver: an error occured, item: %v, err: %v",
					item, err)
			}
		}
	}()
	return out
}

var sniff = false
var client, err = elastic.NewClientFromConfig(&config.Config{
	Sniff: &sniff,
	URL:   cfg.ElasticsearchHosts,
})

// Saving all items.
func Save(item engine.Item, index string) (error) {

	if err != nil {
		return err
	}
	if item.Type == "" {
		return errors.New("type must not be null")
	}
	indexService := client.Index().Index(index).Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
