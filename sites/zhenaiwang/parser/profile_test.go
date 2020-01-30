package parser

import (
	"github.com/lihs-learning/go-crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.html")
	if err != nil {
		panic(err)
	}

	actualProfileResults := ParseProfile(contents)

	if len(actualProfileResults.Items) != 1 {
		t.Errorf("actualProfileResults.Items should contain 1 element, actual contain %d element:\n %v",
			len(actualProfileResults.Items), actualProfileResults.Items)
	}

	profile := actualProfileResults.Items[0].(model.Profile)
	profile.Name = "安静的雪"

	if profile != profileTestCase {
		t.Errorf("expected: %v, actual: %v", profileTestCase, profile)
	}
}
