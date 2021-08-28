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

	if profileExtraInfo.ID != actualParseProfileResults.Items[0].ID {
		t.Errorf("parse result err: except profile id \"%s\", but got \"%s\"",
			profileExtraInfo.ID,
			actualParseProfileResults.Items[0].ID)
	}

	if profileExtraInfo.URL != actualParseProfileResults.Items[0].URL {
		t.Errorf("parse result err: except profile id \"%s\", but got \"%s\"",
			profileExtraInfo.URL,
			actualParseProfileResults.Items[0].URL)
	}

	actualProfile := actualParseProfileResults.Items[0].Payload.(model.Profile)

	if diff := cmp.Diff(actualProfile, exceptProfile); diff != "" {
		t.Errorf("profile mismatch (-excepted +actual):\n%s", diff)
	}
}
