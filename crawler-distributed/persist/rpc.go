package persist

import (
	"polaris/crawler/engine"
	"polaris/crawler/persist"
)

type ItemSaverService struct {
	Index string
}


func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	const index  = "polaris"
	err := persist.Save(item, index)
	if err == nil {
		*result = "ok"
	}
	return err
}
