package parser

import (
	"fmt"
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/qichacha"
	"regexp"
)

var cityRegexp = regexp.MustCompile(
	`<a .*?href="(?P<uri>/firm_(?P<id>[a-f0-9]*)\.html)"[^>]*>(?P<company>.*?)</a>`)

type EnterpriseBrief struct {
	ID   string
	Name string
}

func ParseCity(utf8content []byte) (result engine.ParseResult) {
	enterprisesResult := cityRegexp.FindAllSubmatch(utf8content, -1)

	for _, enterprise := range enterprisesResult {
		enterpriseBrief := EnterpriseBrief{
			ID:   string(enterprise[2]),
			Name: string(filterChinese(enterprise[3])),
		}
		result.Items = append(result.Items,
			enterpriseBrief.Name)
		result.Requests = append(result.Requests, engine.Request{
			URL: fmt.Sprintf("%s%s", qichacha.RootUrl, string(enterprise[1])),
			ParserFunc: func(utf8content []byte) (parseResult engine.ParseResult) {
				return ParseEnterprise(utf8content, enterpriseBrief)
			},
		})
	}

	return
}
