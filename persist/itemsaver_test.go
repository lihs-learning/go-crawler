package persist

import (
	"context"
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/olivere/elastic/v7"
	"testing"

	"github.com/lihs-learning/go-crawler/model"
)

func TestItemSaver(t *testing.T) {
	exceptItem := engine.Item{
		ID:  "8020950579916830882",
		URL: "http://localhost:8080/mock/album.zhenai.com/u/8020950579916830882",
		Payload: model.Profile{
			Name:          "猖狂小丸子",
			Gender:        "女",
			Age:           21,
			Height:        80,
			Weight:        77,
			Income:        "3001-5000元",
			Marriage:      "离异",
			Education:     "初中",
			Occupation:    "程序员",
			Residence:     "青岛市",
			Constellation: "处女座",
			House:         "租房",
			Car:           "有车",
		},
	}

	// TODO: try to start a ES container in docker
	saver, err := NewItemSaver()
	if err != nil {
		panic(err)
	}
	defer func(saver *ItemSaver) {
		err := saver.Close()
		if err != nil {
			panic(err)
		}
	}(saver)

	err = saver.Save(exceptItem)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	res, err := client.Get().
		Index("dating_profile").
		Id(exceptItem.ID).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Log(res)
}
