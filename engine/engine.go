package engine

import (
	"log"
	"time"

	"github.com/lihs-learning/go-crawler/fetcher"
)

func Run(seeds ...Request) {
	requests := make([]Request, 0, len(seeds))
	for _, request := range seeds {
		requests = append(requests, request)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(request.URL)
		if err != nil {
			log.Printf("Fetcher: error in fetching url %s %v", request.URL, err)
			continue
		}

		parseResult := request.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)
		printItems(parseResult.Items)
		time.Sleep(300 * time.Millisecond)
	}
}

func printItems(items []interface{}) {
	for _, item := range items {
		log.Printf("Got item: %v", item)
	}
}
