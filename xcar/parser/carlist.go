package parser

import (
	"fmt"
	"github.com/lihs-learning/go-crawler/xcar"
	"regexp"

	"github.com/lihs-learning/go-crawler/engine"
)

var carModelRe = regexp.MustCompile(`<a .*?href="(?P<modelURI>/\d+/)" .*?class="list_img"[^>]*>`)
var carListRe = regexp.MustCompile(`<a .*?href="//newcar.xcar.com.cn(?P<listURI>/car/[\d+-]+\d+/)"`)

func ParseCarList(utf8Contents []byte) (result engine.ParseResult) {
	modelMatches := carModelRe.FindAllSubmatch(utf8Contents, -1)
	for _, model := range modelMatches {
		result.Requests = append(result.Requests, engine.Request{
			URL: fmt.Sprintf("%s%s", xcar.RootUrl, model[1]),
			ParserFunc: ParseCarModel,
		})
	}

	listMatches := carListRe.FindAllSubmatch(utf8Contents, -1)
	for _, list := range listMatches {
		result.Requests = append(
			result.Requests, engine.Request{
				URL: fmt.Sprintf("%s%s", xcar.RootUrl, list[1]),
				ParserFunc: ParseCarList,
			})
	}

	return result
}
