package main

import (
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/qichacha/parser"
	"github.com/lihs-learning/go-crawler/scheduler"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	concurrentEngine.Run(engine.Request{
		URL:        "https://www.qichacha.com/g_AH.html",
		ParserFunc: parser.ParseProvinceList,
	})
}
