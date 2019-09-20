package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	cityContent, err := ioutil.ReadFile("./city_test_data.html")

	if err != nil {
		panic(err)
	}

	actualParseCityResult := ParseCity(cityContent)

	if len(actualParseCityResult.Items) != EnterprisesSize {
		t.Fatalf("ParseCity result should have %d requests; but had %d",
			EnterprisesSize, len(actualParseCityResult.Items))
	}

	//for i, item := range actualParseCityResult.Items {
	//	fmt.Println(item)
	//	fmt.Println(actualParseCityResult.Requests[i].URL)
	//}

	for i, exceptedEnterprise := range enterprisesTests {
		actualEnterpriseName := actualParseCityResult.Items[i]
		actualEnterPriseURL := actualParseCityResult.Requests[i].URL

		if exceptedEnterprise.Name != actualEnterpriseName {
			t.Errorf("company name expected: \"%s\", actual: \"%s\"",
				exceptedEnterprise.Name, actualEnterpriseName)
		}
		if exceptedEnterprise.URL != actualEnterPriseURL {
			t.Errorf("company url expected: \"%s\", actual: \"%s\"",
				exceptedEnterprise.URL, actualEnterPriseURL)
		}
	}

}
