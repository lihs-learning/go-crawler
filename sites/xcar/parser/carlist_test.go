package parser

import (
	"io/ioutil"
	"regexp"
	"testing"
)

var exceptModelURLs = []string{
	"https://newcar.xcar.com.cn/3428/",
	"https://newcar.xcar.com.cn/3796/",
	"https://newcar.xcar.com.cn/4063/",
	"https://newcar.xcar.com.cn/189/",
	"https://newcar.xcar.com.cn/957/",
	"https://newcar.xcar.com.cn/3866/",
	"https://newcar.xcar.com.cn/3406/",
	"https://newcar.xcar.com.cn/4086/",
	"https://newcar.xcar.com.cn/22/",
	"https://newcar.xcar.com.cn/4116/",
	"https://newcar.xcar.com.cn/1472/",
	"https://newcar.xcar.com.cn/3102/",
	"https://newcar.xcar.com.cn/2823/",
	"https://newcar.xcar.com.cn/3574/",
	"https://newcar.xcar.com.cn/283/",
	"https://newcar.xcar.com.cn/4398/",
	"https://newcar.xcar.com.cn/3900/",
	"https://newcar.xcar.com.cn/1384/",
	"https://newcar.xcar.com.cn/4913/",
	"https://newcar.xcar.com.cn/1406/",
}

var exceptListURLs = []string{
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-10-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-9-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-1-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-2-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-3-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-4-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-5-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-7-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-15-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-16-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-17-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-18-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-19-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-6-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-8-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-11-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-13-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-12-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-14-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-4-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-3-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-7-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-8-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-12-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-13-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-3-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-4-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-2-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-21-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-23-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-22-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-6-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-0-3-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-0-2-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-0-1-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-2-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-1-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-1-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-2-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-3-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-4-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-5-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-6-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-2-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-3-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-4-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-5-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-2-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-4-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-5-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-1-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-3-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-6-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-7-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-8-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-9-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-1-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-2-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-3-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-4-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-5-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-6-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-7-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-8-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-7-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-6-0-0-0-0-1/",
}

func TestParseCarList(t *testing.T) {
	listURLRe := regexp.MustCompile(`/car/[\d+-]+\d+/`)

	contents, err := ioutil.ReadFile("./carlist_test_data.html")
	if err != nil {
		panic(err)
	}

	actualParseCarListResult := ParseCarList(contents)

	var actualModelURLs, actualListURLs []string
	for _, request := range actualParseCarListResult.Requests {
		if listURLRe.MatchString(request.URL) {
			actualListURLs = append(actualListURLs, request.URL)
		} else {
			actualModelURLs = append(actualModelURLs, request.URL)
		}
	}

	if len(exceptModelURLs) != len(actualModelURLs) {
		t.Errorf("except list length: %d, but got actual: %d",
			len(exceptModelURLs), len(actualModelURLs))
	}
	for i := range actualModelURLs {
		if i >= len(exceptModelURLs) {
			t.Errorf("model url expected: \"%T\", actual: \"%s\"",
				nil, actualModelURLs[i])
		} else if exceptModelURLs[i] != actualModelURLs[i] {
			t.Errorf("model url expected: \"%s\", actual: \"%s\"",
				exceptModelURLs[i], actualModelURLs[i])
		}
	}

	if len(exceptListURLs) != len(actualListURLs) {
		t.Errorf("except list length: %d, but got actual: %d",
			len(exceptListURLs), len(actualListURLs))
	}
	for i := range actualListURLs {
		if i >= len(exceptListURLs) {
			t.Errorf("list url expected: \"%T\", actual: \"%s\"",
				nil, actualListURLs[i])
		} else if exceptListURLs[i] != actualListURLs[i] {
			t.Errorf("list url expected: \"%s\", actual: \"%s\"",
				exceptListURLs[i], actualListURLs[i])
		}
	}
}
