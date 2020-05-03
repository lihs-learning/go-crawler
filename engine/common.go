package engine

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v7"

	"github.com/lihs-learning/go-crawler/fetcher"
	"github.com/lihs-learning/go-crawler/redisc"
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

func isDuplicate(url string) bool {
	cmd := redis.NewStringCmd("CF.EXISTS", "crawler", url)
	err := redisc.DB.Process(cmd)
	if err != nil {
		log.Println(err)
	}
	result, err := cmd.Result()
	if err != nil {
		fmt.Println(err)
	}
	duplicated := result == "1"
	if !duplicated {
		cmd := redis.NewStringCmd("CF.ADD", "crawler", url)
		err := redisc.DB.Process(cmd)
		if err != nil {
			fmt.Println(err)
		}
	}
	return duplicated
}
