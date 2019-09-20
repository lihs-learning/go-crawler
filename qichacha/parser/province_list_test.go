package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseProvince(t *testing.T) {
	contents, err := ioutil.ReadFile("./province_test_data.html")

	if err != nil {
		panic(err)
	}

	actualProvinceResult := ParseProvinceList(contents)

	if len(actualProvinceResult.Items) != ProvinceListSize {
		t.Errorf("province result should have %d requests; but had %d",
			ProvinceListSize, len(actualProvinceResult.Requests))
	}

	for i, exceptedProvince := range provinceTests {
		actualProvinceName := actualProvinceResult.Items[i]
		actualProvinceUrl := actualProvinceResult.Requests[i].URL

		if exceptedProvince.Name != actualProvinceName {
			t.Errorf("city name expected: \"%s\", actual: \"%s\"",
				exceptedProvince.Name, actualProvinceName)
		}
		if exceptedProvince.URL != actualProvinceUrl {
			t.Errorf("city url expected: \"%s\", actual: \"%s\"",
				exceptedProvince.URL, actualProvinceUrl)
		}
	}
}
