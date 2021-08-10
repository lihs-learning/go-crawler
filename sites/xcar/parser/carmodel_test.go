package parser

import (
	"io/ioutil"
	"testing"
)

var exceptDetailURLs = []string{
	"https://newcar.xcar.com.cn/m58992/",
	"https://newcar.xcar.com.cn/m58993/",
	"https://newcar.xcar.com.cn/m53617/",
	"https://newcar.xcar.com.cn/m53693/",
	"https://newcar.xcar.com.cn/m44730/",
	"https://newcar.xcar.com.cn/m44732/",
}

func TestParseCarModel(t *testing.T) {
	contents, err := ioutil.ReadFile("./carmodel_test_data.html")
	if err != nil {
		panic(err)
	}

	actualParseCarModelResult := ParseCarModel(contents)

	if len(exceptDetailURLs) != len(actualParseCarModelResult.Requests) {
		t.Errorf("except detail length: %d, but got actual: %d",
			len(exceptDetailURLs), len(actualParseCarModelResult.Requests))
	}

	for i := range actualParseCarModelResult.Requests {
		if i >= len(exceptDetailURLs) {
			t.Errorf("detail url expected: %T, actual: \"%s\"",
				nil, actualParseCarModelResult.Requests[i].URL)
		} else if exceptDetailURLs[i] != actualParseCarModelResult.Requests[i].URL {
			t.Errorf("detail url expected: \"%s\", actual: \"%s\"",
				exceptDetailURLs[i], actualParseCarModelResult.Requests[i].URL)
		}
	}
}
