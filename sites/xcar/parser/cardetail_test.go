package parser

import (
	"bufio"
	"github.com/google/go-cmp/cmp"
	"github.com/lihs-learning/go-crawler/model"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func determineEncoding(bufReader *bufio.Reader) encoding.Encoding {
	bytes, err := bufReader.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func TestParseCarDetail(t *testing.T) {
	f, err := os.Open("./cardetail_test_data.html")
	if err != nil {
		panic(err)
	}

	bodyReader := bufio.NewReader(f)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	contents, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	actualParseCarDetailResult := ParseCarDetail(contents, "https://newcar.xcar.com.cn/m49989/")

	if len(actualParseCarDetailResult.Items) != 1 {
		t.Errorf("car detail result shoud have len=1 but got len=%d",
			len(actualParseCarDetailResult.Items))
	}

	actualCar, ok := actualParseCarDetailResult.Items[0].(model.Car)
	if !ok {
		t.Errorf("type err except result item type %T", model.Car{})
	}

	exceptedCar := model.Car{
		Name:         "零跑S012021款460 酷玩版",
		Price:        130900,
		ImageURL:     "https://img1.xcarimg.com/PicLib/s/s13985_300.jpg",
		Size:         "4075×1760×1380mm",
		Fuel:         0,
		Transmission: "电动车单速变速箱",
		Engine:       "125kW(电动机)",
		Displacement: 0,
		MaxSpeed:     135,
		Acceleration: 0,
		SourceURL:    "https://newcar.xcar.com.cn/m49989/",
	}

	if diff := cmp.Diff(exceptedCar, actualCar); diff != "" {
		t.Errorf("car mismatch (-excepted +actual):\n%s", diff)
	}
}
