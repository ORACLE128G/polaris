package persist

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"polaris/crawler/engine"
	"polaris/crawler/persist"
)

type ItemSaverService struct {
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	const index = "polaris"
	err := persist.Save(item, index)
	if err == nil {
		fmt.Printf("saving item#%v \n", item.Url)
		*result = "ok"
	} else {
		log.Error("saving item#%v fail", item)
		*result = "failed"
	}
	return err
}
