package parser

import (
	"fmt"
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/qichacha"
	"regexp"
)

var cityRegexp = regexp.MustCompile(
	`<a .*?href="(?P<uri>/firm_[a-f0-9]*\.html)"[^>]*>(?P<company>.*?)</a>`)

func ParseCity(utf8content []byte) (result engine.ParseResult) {
	companyResult := cityRegexp.FindAllSubmatch(utf8content, -1)

	for _, company := range companyResult {
		result.Items = append(result.Items,
			string(filterChinese(company[2])))
		result.Requests = append(result.Requests, engine.Request{
			URL:        fmt.Sprintf("%s%s", qichacha.RootUrl, string(company[1])),
			ParserFunc: nil,
		})
	}

	return
}
