package main

import (
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/qichacha/parser"
)

func main() {
	engine.Run(engine.Request{
		URL:        "https://www.qichacha.com/g_AH.html",
		ParserFunc: parser.ParseProvinceList,
	})
}
