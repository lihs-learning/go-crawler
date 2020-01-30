package main

import (
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/scheduler"

	"github.com/lihs-learning/go-crawler/sites/xcar/parser"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	concurrentEngine.Run(engine.Request{
		URL:        "https://newcar.xcar.com.cn/",
		ParserFunc: parser.ParseCarList,
	})
}
