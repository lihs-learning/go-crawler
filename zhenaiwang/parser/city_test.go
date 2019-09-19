package parser

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("./city_test_data.html")

	if err != nil {
		panic(err)
	}

	actualUserResults := ParseCity(contents)

	if len(actualUserResults.Requests) != len(cityTestCases) {
		t.Errorf("city result should have %d requests; but had %d\n%v",
			len(cityTestCases), len(actualUserResults.Requests), actualUserResults)
		return
	}

	for i, expectedUserResult := range cityTestCases {
		actualCityName := actualUserResults.Items[i]
		actualCityURL := actualUserResults.Requests[i].URL
		if expectedUserResult.Nick != actualCityName {
			t.Errorf("city name expected: %s, actual: %s",
				expectedUserResult.Nick, actualCityName)
		}
		if expectedUserResult.URL != actualCityURL {
			t.Errorf("city url expected: %s, actual: %s",
				expectedUserResult.URL, actualCityURL)
		}
	}
}

func TestProfileRe(t *testing.T) {
	profileRegexp := regexp.MustCompile(profileReofileRe)
	for _, test := range profileRegExpTestCases {
		groupNamesMap := make(map[string]int)

		for groupIndex, groupName := range profileRegexp.SubexpNames() {
			groupNamesMap[groupName] = groupIndex
		}
		profile := profileRegexp.FindSubmatch([]byte(test.htmlTagString))
		if len(profile) == 0 && (test.exceptedProfile != profileResult{}) {
			t.Errorf("test profile string: %s\nresult expected: %v, actual: %v",
				test.htmlTagString, test.exceptedProfile, profileResult{})
			return
		}
		actualProfileResult := profileResult{
			URL:  string(profile[groupNamesMap["link"]]),
			Nick: string(profile[groupNamesMap["name"]]),
		}
		if actualProfileResult != test.exceptedProfile {
			t.Errorf("test profile string: %s\nresult expected: %v, actual: %v",
				test.htmlTagString, test.exceptedProfile, actualProfileResult)
		}
	}
}
