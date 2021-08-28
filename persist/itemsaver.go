package persist

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"

	"github.com/lihs-learning/go-crawler/engine"
)

type ItemSaver struct {
	IsClosed bool
	Out chan engine.Item
}

func NewItemSaver() (itemSaver *ItemSaver, err error) {
	itemSaver = &ItemSaver{
		Out: make(chan engine.Item),
	}
	client, err := elastic.NewClient(
		// must turn off when elasticsearch is running in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	go func() {
		for item := range itemSaver.Out {
			indexService := client.Index().
				Index("dating_profile")
			if item.ID != "" {
				indexService = indexService.Id(item.ID)
			}
			_, err := indexService.
				BodyJson(item).
				Do(context.Background())
			if err != nil {
				log.Printf("ItemSaver err: %s, with item: %v", err, item)
			}
		}
	}()
	return
}

func (saver *ItemSaver) Save(item engine.Item) (err error) {
	if saver.IsClosed {
		return fmt.Errorf("saver already closed")
	}
	saver.Out <- item
	return
}

func (saver *ItemSaver) Close() (err error) {
	if saver.IsClosed {
		return fmt.Errorf("saver already closed")
	}
	close(saver.Out)
	saver.IsClosed = true
	return
}
