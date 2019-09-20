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

	if len(actualParseCityResult.Requests) != CompanySize {
		t.Fatalf("ParseCity result should have %d requests; but had %d",
			ProvinceListSize, len(actualParseCityResult.Requests))
	}

	//for i, item := range actualParseCityResult.Items {
	//	fmt.Println(item)
	//	fmt.Println(actualParseCityResult.Requests[i].URL)
	//}

	for i, exceptedCompany := range companyTests {
		actualCompanyName := actualParseCityResult.Items[i]
		actualCompanyURL := actualParseCityResult.Requests[i].URL

		if exceptedCompany.Name != actualCompanyName {
			t.Errorf("company name expected: \"%s\", actual: \"%s\"",
				exceptedCompany.Name, actualCompanyName)
		}
		if exceptedCompany.URL != actualCompanyURL {
			t.Errorf("company url expected: \"%s\", actual: \"%s\"",
				exceptedCompany.URL, actualCompanyURL)
		}
	}

}
