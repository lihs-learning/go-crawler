package engine

import (
	"log"
)

type SimpleEngine struct{
	ItemChan chan interface{}
}

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
		for _, item := range parseResult.Items {
			e.ItemChan <- item
		}
	}
}
