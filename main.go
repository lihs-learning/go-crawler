package main

import (
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/zhenaiwang/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}











