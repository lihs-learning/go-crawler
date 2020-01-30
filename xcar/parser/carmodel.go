package parser

import (
	"fmt"
	"regexp"

	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/xcar"
)

var carDetailRe = regexp.MustCompile(`<a .*?href="(?P<modelURI>/m\d+/)"`)

func ParseCarModel(utf8Contents []byte) (result engine.ParseResult) {
	detailMatches := carDetailRe.FindAllSubmatch(utf8Contents, -1)
	for _, detail := range detailMatches {
		result.Requests = append(
			result.Requests, engine.Request{
				URL: fmt.Sprintf("%s%s", xcar.RootUrl, detail[1]),
				ParserFunc: ParseCarDetail,
			})
	}

	return result
}
