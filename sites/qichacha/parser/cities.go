package parser

import (
	"fmt"
	"regexp"

	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/sites/qichacha"
)

var cityListRegexp = regexp.MustCompile(
	`<a .*?href="/g_(?P<prov>[A-Z]*)_(?P<city>\d{6})\.html"[^>]*>(?P<name>[^<]*)</a>`)

func ParseCityList(utf8content []byte) (result engine.ParseResult) {
	allCity := cityListRegexp.FindAllSubmatch(utf8content, -1)
	for _, city := range allCity {
		//result.Items = append(result.Items, string(city[3]))
		result.Requests = append(result.Requests,
			engine.Request{
				URL: fmt.Sprintf("%s/gongsi_area.html?prov=%s&city=%s&p=1",
					qichacha.RootUrl, city[1], city[2]),
				ParserFunc: ParseCity,
			})
	}
	return
}
