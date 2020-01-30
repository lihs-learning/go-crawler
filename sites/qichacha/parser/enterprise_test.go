package parser

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/lihs-learning/go-crawler/model"
)

func TestParseEnterprise(t *testing.T) {
	contents, err := ioutil.ReadFile("./enterprise_test_data.html")

	if err != nil {
		panic(err)
	}

	actualParseEnterpriseResult := ParseEnterprise(contents, EnterpriseBriefDemo)

	if len(actualParseEnterpriseResult.Items) != 1 {
		t.Errorf("province result should have %d requests; but had %d",
			1, len(actualParseEnterpriseResult.Items))
	}

	actualEnterprise := actualParseEnterpriseResult.Items[0].(model.Enterprise)

	if diff := cmp.Diff(enterpriseTest, actualEnterprise); diff != "" {
		t.Errorf("enterprise mismatch (-excepted +actual):\n%s", diff)
	}
}
