package parser

const CityListSize = 470

var cityListTestCases = []struct {
	CityName string
	CityURL  string
}{
	{
		CityName: "阿坝",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
	}, {
		CityName: "阿克苏",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
	}, {
		CityName: "阿拉善盟",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}, {
		CityName: "阿勒泰",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/aletai",
	}, {
		CityName: "阿里",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/ali",
	}, {
		CityName: "安徽",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/anhui",
	}, {
		CityName: "安康",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/ankang",
	}, {
		CityName: "安庆",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/anqing",
	}, {
		CityName: "鞍山",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/anshan",
	}, {
		CityName: "安顺",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/anshun",
	}, {
		CityName: "安阳",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/anyang",
	}, {
		CityName: "澳门",
		CityURL:  "http://localhost:8080/mock/www.zhenai.com/zhenghun/aomen",
	},
}
