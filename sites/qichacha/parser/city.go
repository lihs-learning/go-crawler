package parser

import (
	"fmt"
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/sites/qichacha"
	"regexp"
)

var cityRegexp = regexp.MustCompile(
	`<a .*?href="(?P<uri>/firm_(?P<id>[a-f0-9]*)\.html)"[^>]*>(?P<company>.*?)</a>`)

func ParseCity(utf8content []byte) (result engine.ParseResult) {
	enterprisesResult := cityRegexp.FindAllSubmatch(utf8content, -1)

	for _, enterprise := range enterprisesResult {
		enterpriseURL := fmt.Sprintf("%s%s", qichacha.RootUrl, string(enterprise[1]))
		enterpriseBrief := EnterpriseBrief{
			ID:   string(enterprise[2]),
			URL: enterpriseURL,
			Name: string(filterChinese(enterprise[3])),
		}
		result.Requests = append(result.Requests, engine.Request{
			URL: enterpriseURL,
			ParserFunc: func(utf8content []byte) (parseResult engine.ParseResult) {
				return ParseEnterprise(utf8content, enterpriseBrief)
			},
		})
	}

	return
}
