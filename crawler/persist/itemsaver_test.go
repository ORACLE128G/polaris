package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"polaris/crawler/engine"
	"polaris/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/110591409",
		Type: "zhenai",
		Id:   "110591409",
		PayLoad: model.Profile{
			Name:       "征婚",
			Age:        30,
			Height:     180,
			Gender:     "",
			Weight:     120,
			Income:     "",
			Marriage:   "未婚",
			Education:  "大专",
			Occupation: "",
			Hukou:      "四川",
			Xinzuo:     "白羊",
			House:      "未购房",
			Car:        "未购车",
		},
	}

	// Save expected item.
	err := Save(expected, "polaris")

	if err != nil {
		panic(err)
	}

	// Deserialization data by id.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// Fetch saved item.
	r, err := client.Get().
		Index("polaris").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	json.Unmarshal(*r.Source, &actual)

	profile, _ := model.FromJson(actual.PayLoad)
	actual.PayLoad = profile
	// Verify result.
	if actual != expected {
		t.Errorf("Testing error: expected: %v, actual: %v",
			expected, actual)
	}
}
