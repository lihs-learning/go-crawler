package parser

import (
	"fmt"
	xcar2 "github.com/lihs-learning/go-crawler/sites/xcar"
	"regexp"

	"github.com/lihs-learning/go-crawler/engine"
)

var carDetailRe = regexp.MustCompile(`<a .*?href="(?P<modelURI>/m\d+/)"`)

func ParseCarModel(utf8Contents []byte) (result engine.ParseResult) {
	detailMatches := carDetailRe.FindAllSubmatch(utf8Contents, -1)
	for _, detail := range detailMatches {
		result.Requests = append(
			result.Requests, engine.Request{
				URL:        fmt.Sprintf("%s%s", xcar2.RootUrl, detail[1]),
				ParserFunc: ParseCarDetail,
			})
	}

	return result
}