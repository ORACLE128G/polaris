package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"polaris/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:   "Polaris",
		Age:    18,
		Height: 175,
	}
	id, err := save(expected)

	if err != nil {
		panic(err)
	}

	// Deserialization data by id.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	r , err:= client.Get().
		Index("polaris").
		Type("zhenai").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual model.Profile
	err = json.Unmarshal(*r.Source, &actual)
	if err != nil {
		panic(err)
	}
	t.Logf("Got source: %+v", actual)
}
