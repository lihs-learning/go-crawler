package engine

import (
	"log"
	"time"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	requests := make([]Request, 0, len(seeds))
	for _, request := range seeds {
		requests = append(requests, request)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parseResult, err := worker(request)
		if err != nil {
			log.Printf("[Worker Error] %s", err)
			continue
		}

		requests = append(requests, parseResult.Requests...)
		printItems(parseResult.Items)
		time.Sleep(300 * time.Millisecond)
	}
}
