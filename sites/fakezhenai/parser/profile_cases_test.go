package parser

import "github.com/lihs-learning/go-crawler/model"

var profileExtraInfo = ProfileParserExtraInfo{
	ID: "8020950579916830882",
	URL: "http://localhost:8080/mock/album.zhenai.com/u/8020950579916830882",
	Name:   "猖狂小丸子",
}

var exceptProfile = model.Profile{
	Name:          "猖狂小丸子",
	Gender:        "女",
	Age:           21,
	Height:        80,
	Weight:        77,
	Income:        "3001-5000元",
	Marriage:      "离异",
	Education:     "初中",
	Occupation:    "程序员",
	Residence:     "青岛市",
	Constellation: "处女座",
	House:         "租房",
	Car:           "有车",
}
