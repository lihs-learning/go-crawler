package engine

import "encoding/json"

type Request struct {
	URL        string
	ParserFunc func(utf8Content []byte) (parseResult ParseResult)
}

type Item struct {
	ID      string      `json:"id,omitempty"`
	URL     string      `json:"url,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

// MarshalJSON TODO: Performance is bad, but work
func (i Item) MarshalJSON() ([]byte, error) {
	payloadJSON, err := json.Marshal(i.Payload)
	if err != nil {
		return nil, err
	}
	tmp := make(map[string]interface{})
	tmp["url"] = i.URL
	err = json.Unmarshal(payloadJSON, &tmp)
	if err != nil {
		return nil, err
	}
	return json.Marshal(tmp)
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

func NilParser(_ []byte) ParseResult {
	return ParseResult{}
}

//14-6 测试CityListParser 02:13

type Persist interface {
	Save(item Item) (err error)
}
