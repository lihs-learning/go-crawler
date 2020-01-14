package engine

import (
	"fmt"
	"log"

	"github.com/lihs-learning/go-crawler/fetcher"
)

func worker(request Request) (parseResult ParseResult, err error) {
	body, fetchErr := fetcher.Fetch(request.URL)
	if fetchErr != nil {
		err = fmt.Errorf("in fetching url %s, got \"%v\"", request.URL, fetchErr)
		return ParseResult{}, err
	}

	parseResult = request.ParserFunc(body)
	return
}

func printItems(items []interface{}) {
	for _, item := range items {
		log.Printf("Got item: %v", item)
	}
}
