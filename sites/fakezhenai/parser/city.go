package parser

import (
	"github.com/lihs-learning/go-crawler/engine"
	"regexp"
)

var userRegexp = regexp.MustCompile(
	`<a .*?href="(?P<link>https?://localhost:8080/mock/album\.zhenai\.com/u/(?P<uid>\d+))"[^>]*>\s*(?P<name>[^\s<]+(?:\s+[^\s<]+)*)\s*</a>`)
var cityPageRegexp = regexp.MustCompile(
	`<a [.\s]*href="(?P<link>https?://localhost:8080/mock/www\.zhenai\.com/zhenghun/[0-9a-zA-Z]+/\d+)"[^>]*>(?P<page>[^<]+)</a>`)

func ParseCity(utf8Content []byte) (result engine.ParseResult) {
	userGroupNamesMap := make(map[string]int)

	for groupIndex, groupName := range userRegexp.SubexpNames() {
		userGroupNamesMap[groupName] = groupIndex
	}
	allUsers := userRegexp.FindAllSubmatch(utf8Content, -1)
	for _, match := range allUsers {
		userName := string(match[userGroupNamesMap["name"]])
		userID := string(match[userGroupNamesMap["uid"]])
		userURL := string(match[userGroupNamesMap["link"]])
		result.Requests = append(result.Requests,
			engine.Request{
				URL: userURL,
				ParserFunc: func(utf8Content []byte) (parseResult engine.ParseResult) {
					return ParseProfile(utf8Content, ProfileParserExtraInfo{
						ID: userID,
						URL: userURL,
						Name:   userName,
					})
				},
			})
	}

	cityPageGroupNamesMap := make(map[string]int)
	for groupIndex, groupName := range cityPageRegexp.SubexpNames() {
		cityPageGroupNamesMap[groupName] = groupIndex
	}
	allCityPages := cityPageRegexp.FindAllSubmatch(utf8Content, -1)
	for _, match := range allCityPages {
		cityPageURL := string(match[cityPageGroupNamesMap["link"]])
		result.Requests = append(result.Requests,
			engine.Request{
				URL: cityPageURL,
				ParserFunc: ParseCity,
			})
	}

	return result
}
