package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("./cities_test_data.html")

	if err != nil {
		panic(err)
	}

	actualParseCitiesResult := ParseCityList(contents)

	if len(actualParseCitiesResult.Items) != CitiesSize {
		t.Errorf("city result should have %d requests; but had %d",
			CitiesSize, len(actualParseCitiesResult.Items))
	}

	//for i, item := range actualParseCitiesResult.Items {
	//	fmt.Println(item)
	//	fmt.Println(actualParseCitiesResult.Requests[i].URL)
	//}

	for i, exceptedCity := range citiesTests {
		actualCityName := actualParseCitiesResult.Items[i]
		actualCityURL := actualParseCitiesResult.Requests[i].URL

		if exceptedCity.Name != actualCityName {
			t.Errorf("city name expected: \"%s\", actual: \"%s\"",
				exceptedCity.Name, actualCityName)
		}
		if exceptedCity.URL != actualCityURL {
			t.Errorf("city url expected: \"%s\", actual: \"%s\"",
				exceptedCity.URL, actualCityURL)
		}
	}
}
