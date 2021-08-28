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

func isDuplicate(url string) bool {
	cmd := redis.NewStringCmd("CF.ADDNX", "crawler", url)
	err := redisc.DB.Process(cmd)
	if err != nil {
		log.Printf("isDuplicate::redisc.DB.Process err: %v, with cmd: \"%s\"", err, cmd)
	}
	result, err := cmd.Result()
	if err != nil {
		log.Printf("isDuplicate::cmd.Result err: %v", err)
	}
	return result == "0"
}
