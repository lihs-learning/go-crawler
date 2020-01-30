package parser

import (
	"io/ioutil"
	"testing"
)

var exceptDetailURLs = []string{
	"https://newcar.xcar.com.cn/m49989/",
	"https://newcar.xcar.com.cn/m50314/",
	"https://newcar.xcar.com.cn/m50315/",
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
		if exceptDetailURLs[i] != actualParseCarModelResult.Requests[i].URL {
			t.Errorf("model url expected: \"%s\", actual: \"%s\"",
				exceptDetailURLs[i], actualParseCarModelResult.Requests[i].URL)
		}
	}
}

