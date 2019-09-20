package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("./city_list_test_data.html")

	if err != nil {
		panic(err)
	}

	actualCityResult := ParseCityList(contents)

	if len(actualCityResult.Items) != CityListSize {
		t.Errorf("city result should have %d requests; but had %d",
			ProvinceListSize, len(actualCityResult.Requests))
	}

	//for i, item := range actualCityResult.Items {
	//	fmt.Println(item)
	//	fmt.Println(actualCityResult.Requests[i].URL)
	//}

	for i, exceptedCity := range cityTests {
		actualCityName := actualCityResult.Items[i]
		actualCityUrl := actualCityResult.Requests[i].URL

		if exceptedCity.Name != actualCityName {
			t.Errorf("city name expected: \"%s\", actual: \"%s\"",
				exceptedCity.Name, actualCityName)
		}
		if exceptedCity.URL != actualCityUrl {
			t.Errorf("city url expected: \"%s\", actual: \"%s\"",
				exceptedCity.URL, actualCityUrl)
		}
	}
}