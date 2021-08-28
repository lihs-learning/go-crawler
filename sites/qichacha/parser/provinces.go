package parser

import (
	"fmt"
	"regexp"

	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/sites/qichacha"
)

var provinceRegexp = regexp.MustCompile(
	`<a .*?href="(?P<uri>/g_[A-Z]*\.html)"[^>]*>(?P<name>[^<]*)</a>`)

func ParseProvinceList(utf8Content []byte) (result engine.ParseResult) {
	allProvince := provinceRegexp.FindAllSubmatch(utf8Content, -1)
	for _, province := range allProvince {
		//result.Items = append(result.Items, string(filterChinese(province[2])))
		result.Requests = append(result.Requests,
			engine.Request{
				URL:        fmt.Sprintf("%s%s", qichacha.RootUrl, province[1]),
				ParserFunc: ParseCityList,
			})
	}
	return
}
