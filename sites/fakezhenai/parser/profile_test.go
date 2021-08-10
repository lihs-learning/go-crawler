package parser

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lihs-learning/go-crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.html")
	if err != nil {
		panic(err)
	}

	actualParseProfileResults := ParseProfile(contents, profileExtraInfo)

	if len(actualParseProfileResults.Items) != 1 {
		t.Errorf("actualProfileResults.Items should contain 1 element, actual contain %d element:\n %v",
			len(actualParseProfileResults.Items), actualParseProfileResults.Items)
	}

	actualProfile := actualParseProfileResults.Items[0].(model.Profile)

	if diff := cmp.Diff(actualProfile, exceptProfile); diff != "" {
		t.Errorf("profile mismatch (-excepted +actual):\n%s", diff)
	}
}
