package parser

import (
	"github.com/lihs-learning/go-crawler/engine"
	"regexp"
)

const profileRe = `<a .*?href="(?P<link>https?://album\.zhenai\.com/u/\d+)"[^>]*>\s*(?P<name>[^\s<]+(?:\s+[^\s<]+)*)\s*</a>`

var userRegexp = regexp.MustCompile(profileRe)

func ParseCity(utf8Content []byte) (result engine.ParseResult) {
	groupNamesMap := make(map[string]int)

	for groupIndex, groupName := range userRegexp.SubexpNames() {
		groupNamesMap[groupName] = groupIndex
	}
	allUsers := userRegexp.FindAllSubmatch(utf8Content, -1)
	for _, match := range allUsers {
		result.Items = append(result.Items,
			string(match[groupNamesMap["name"]]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(match[groupNamesMap["link"]]),
				ParserFunc: engine.NilParser,
			})
	}

	return result
}
