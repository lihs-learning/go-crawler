package parser

import (
	"github.com/lihs-learning/go-crawler/engine"
	"regexp"
)

var userRegexp = regexp.MustCompile(
	`<a .*?href="(?P<link>https?://album\.zhenai\.com/u/\d+)"[^>]*>\s*(?P<name>[^\s<]+(?:\s+[^\s<]+)*)\s*</a>`)

func ParseCity(utf8Content []byte) (result engine.ParseResult) {
	groupNamesMap := make(map[string]int)

	for groupIndex, groupName := range userRegexp.SubexpNames() {
		groupNamesMap[groupName] = groupIndex
	}
	allUsers := userRegexp.FindAllSubmatch(utf8Content, -1)
	for _, match := range allUsers {
		userName := string(match[groupNamesMap["name"]])
		result.Items = append(result.Items, userName)
		result.Requests = append(result.Requests,
			engine.Request{
				URL: string(match[groupNamesMap["link"]]),
				ParserFunc: func(utf8Content []byte) (parseResult engine.ParseResult) {
					return ParseProfile(utf8Content, ProfileParserExtraInfo{
						Name:   userName,
						Gender: userGender,
					})
				},
			})
	}

	return result
}
