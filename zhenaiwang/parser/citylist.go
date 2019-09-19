package parser

import (
	"regexp"

	"github.com/lihs-learning/go-crawler/engine"
)

var cityRegexp = regexp.MustCompile(
`<a [.\s]*href="(?P<link>https?://www\.zhenai\.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>(?P<name>[^<]+)</a>`)

func ParseCityList(utf8Content []byte) (result engine.ParseResult) {
	groupNamesMap := make(map[string]int)

	for groupIndex, groupName := range cityRegexp.SubexpNames() {
		groupNamesMap[groupName] = groupIndex
	}
	allCities := cityRegexp.FindAllSubmatch(utf8Content, -1)
	for _, match := range allCities {
		result.Items = append(result.Items,
			string(match[groupNamesMap["name"]]))
		result.Requests = append(result.Requests,
			engine.Request{
				URL:        string(match[groupNamesMap["link"]]),
				ParserFunc: ParseCity,
			})
	}

	return result
}
