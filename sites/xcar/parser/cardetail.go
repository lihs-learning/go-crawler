package parser

import (
	"fmt"
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/model"
	"regexp"
)

var nameRe = regexp.MustCompile(`<title>【(.*)】报价_图片_参数(.*)</title>`)
var carImageRe = regexp.MustCompile(`<img class="color_car_img_new" src="//([^"]+)"`)
var sizeRe = regexp.MustCompile(`<li.*车身尺寸.*<em>(\d+[^\d]\d+[^\d]\d+mm)`)
var fuelRe = regexp.MustCompile(`<li.*工信部油耗.*<em>(\d+\.\d+)L/100km`)
var transmissionRe = regexp.MustCompile(`<li.*变\s*速\s*箱.*<em>(.+)</em>`)
var engineRe = regexp.MustCompile(`发\s*动\s*机.*\s*.*<.*>(\d+kW[^<]*)<`)
var displacementRe = regexp.MustCompile(`<li.*排.*量.*(\d+\.\d+)L`)
var maxSpeedRe = regexp.MustCompile(`<td.*最高车速\(km/h\).*\s*<td[^>]*>(\d+)</td>`)
var accelRe = regexp.MustCompile(`<td.*0-100加速时间\(s\).*\s*<td[^>]*>(\d+(?:\.\d+))</td>`)
var costRe = regexp.MustCompile(`<div class="cost"><span>裸车价格：(\d*)元</span></div>`)

func ParseCarDetail(utf8contents []byte, srcURL string) (result engine.ParseResult) {
	car := model.Car{
		Name:         extractString(utf8contents, nameRe),
		Price:        extractFloat(utf8contents, costRe),
		ImageURL:     fmt.Sprintf("https://%s", extractString(utf8contents, carImageRe)),
		Size:         extractString(utf8contents, sizeRe),
		Fuel:         extractFloat(utf8contents, fuelRe),
		Transmission: extractString(utf8contents, transmissionRe),
		Engine:       extractString(utf8contents, engineRe),
		Displacement: extractFloat(utf8contents, displacementRe),
		MaxSpeed:     extractFloat(utf8contents, maxSpeedRe),
		Acceleration: extractFloat(utf8contents, accelRe),
		SourceURL:    srcURL,
	}

	result.Items = append(result.Items, car)
	return
}
