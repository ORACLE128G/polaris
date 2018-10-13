package persist

import (
	"polaris/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:   "Polaris",
		Age:    18,
		Height: 175,
	}
	id, _ := save(expected)
	t.Logf("id: %v", id)
	/*id, err := save(expected)
	if err != nil {
		panic(err)
	}*/
}
