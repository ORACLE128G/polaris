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
	const index  = "polaris"
	err := persist.Save(item, index)
	if err == nil {
		fmt.Println("saving item success")
		*result = "ok"
	} else {
		log.Error("saving item fail")
		*result = "failed"
	}
	return err
}
