package engine

type Request struct {
	Url        string
	ParserFunc func(utf8Content []byte) (parseResult ParseResult)
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser(_ []byte) ParseResult {
	return ParseResult{}
}

//14-6 测试CityListParser 02:13
