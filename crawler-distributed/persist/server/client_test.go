package main

import (
	"polaris/crawler-distributed/config"
	"polaris/crawler-distributed/rpc"
	"polaris/crawler/engine"
	"polaris/crawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {

	// start ItemSaverServer
	go serveRpc(config.ItemSaver0Host, "polaris")

	// start ItemSaverClient
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(config.ItemSaver0Host)
	if err != nil {
		panic(err)
	}

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

	result := ""
	err = client.Call("ItemSaverService.Save",
		expected, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %v; error: %v", result, err)
	}
}
