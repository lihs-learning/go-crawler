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
	"https://newcar.xcar.com.cn/63/",
	"https://newcar.xcar.com.cn/3226/",
	"https://newcar.xcar.com.cn/10/",
	"https://newcar.xcar.com.cn/4475/",
	"https://newcar.xcar.com.cn/1468/",
	"https://newcar.xcar.com.cn/306/",
	"https://newcar.xcar.com.cn/4039/",
	"https://newcar.xcar.com.cn/3028/",
	"https://newcar.xcar.com.cn/3563/",
	"https://newcar.xcar.com.cn/2805/",
}

var exceptListURLs = []string{
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-6-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-8-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-11-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-13-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-12-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-14-0-0-0-0-1/",
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
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-4-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-3-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-7-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-8-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-12-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-13-0-0-0-0-0-0-0-0-0-1/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-0-0/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-3-0/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-4-0/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-2-0/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-21-0/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-23-0/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-22-0/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-6-0/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-0-0/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-0-3/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-0-2/",
	"https://newcar.xcar.com.cn/car/0-0-0-0-0-0-0-0-0-0-0-1-0-1/",
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
		if exceptModelURLs[i] != actualModelURLs[i] {
			t.Errorf("model url expected: \"%s\", actual: \"%s\"",
				exceptModelURLs[i], actualModelURLs[i])
		}
	}

	if len(exceptListURLs) != len(actualListURLs) {
		t.Errorf("except list length: %d, but got actual: %d",
			len(exceptListURLs), len(actualListURLs))
	}
	for i := range actualListURLs {
		if exceptListURLs[i] != actualListURLs[i] {
			t.Errorf("model url expected: \"%s\", actual: \"%s\"",
				exceptListURLs[i], actualListURLs[i])
		}
	}
}
