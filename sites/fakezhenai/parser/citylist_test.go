package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("./citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	actualCityResults := ParseCityList(contents)

	if len(actualCityResults.Requests) != CityListSize {
		t.Errorf("city result should have %d requests; but had %d",
			CityListSize, len(actualCityResults.Requests))
	}

	for i, expectedCityResult := range cityListTestCases {
		actualCityURL := actualCityResults.Requests[i].URL
		if expectedCityResult.CityURL != actualCityURL {
			t.Errorf("city url expected: %s, actual: %s",
				expectedCityResult.CityURL, actualCityURL)
		}
	}
}
