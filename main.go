package main

import (
	"log"

	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/persist"
	"github.com/lihs-learning/go-crawler/scheduler"

	"github.com/lihs-learning/go-crawler/sites/fakezhenai/parser"
)

func main() {
	itemSaver, err := persist.NewItemSaver()
	if err != nil {
		log.Printf("create persist error: %s", err)
		return
	}
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		Persist:     itemSaver,
	}
	concurrentEngine.Run(engine.Request{
		URL:        "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
