package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseProvince(t *testing.T) {
	contents, err := ioutil.ReadFile("./provinces_test_data.html")

	if err != nil {
		panic(err)
	}

	actualParseProvincesResult := ParseProvinceList(contents)

	//if len(actualParseProvincesResult.Items) != ProvincesSize {
	//	t.Errorf("province result should have %d requests; but had %d",
	//		ProvincesSize, len(actualParseProvincesResult.Items))
	//}

	for i, exceptedProvince := range provinceTests {
		//actualProvinceName := actualParseProvincesResult.Items[i]
		actualProvinceURL := actualParseProvincesResult.Requests[i].URL

		//if exceptedProvince.Name != actualProvinceName {
		//	t.Errorf("city name expected: \"%s\", actual: \"%s\"",
		//		exceptedProvince.Name, actualProvinceName)
		//}
		if exceptedProvince.URL != actualProvinceURL {
			t.Errorf("city url expected: \"%s\", actual: \"%s\"",
				exceptedProvince.URL, actualProvinceURL)
		}
	}
}
