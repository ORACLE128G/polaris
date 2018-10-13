package persist

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"log"
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
			err := save(item)
			if err != nil {
				log.Printf("Item Saver: an error occured, item: %v, err: %v",
					item, err)
			}
		}
	}()
	return out
}

var client, _ = elastic.NewClient(
	elastic.SetSniff(false))
// Saving all items.
func save(item engine.Item) (err error) {
	if err != nil {
		return err
	}
	if item.Type == "" {
		return errors.New("type must not be null")
	}
	indexService := client.Index().Index("polaris").Type(item.Type).
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
