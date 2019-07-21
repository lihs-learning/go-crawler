package parser

import "fmt"

type profileResult struct {
	Nick string
	URL  string
}

var cityTestCases = []profileResult{
	{
		Nick: "静待一人",
		URL:  "http://album.zhenai.com/u/1069129647",
	}, {
		Nick: "劣酒灼心",
		URL:  "http://album.zhenai.com/u/1212756553",
	}, {
		Nick: "扎西",
		URL:  "http://album.zhenai.com/u/1722799091",
	}, {
		Nick: "绘天",
		URL:  "http://album.zhenai.com/u/1967129012",
	}, {
		Nick: "一场电影而已",
		URL:  "http://album.zhenai.com/u/1321512066",
	}, {
		Nick: "半生緣",
		URL:  "http://album.zhenai.com/u/1480921860",
	}, {
		Nick: "脚踏实地",
		URL:  "http://album.zhenai.com/u/1396372515",
	}, {
		Nick: "不想",
		URL:  "http://album.zhenai.com/u/1426975040",
	}, {
		Nick: "我们结婚吧",
		URL:  "http://album.zhenai.com/u/1042163302",
	}, {
		Nick: "lucky傩",
		URL:  "http://album.zhenai.com/u/1877553985",
	}, {
		Nick: "忘忧",
		URL:  "http://album.zhenai.com/u/1587689291",
	}, {
		Nick: "亦梦亦真",
		URL:  "http://album.zhenai.com/u/1928969584",
	}, {
		Nick: "zouwei",
		URL:  "http://album.zhenai.com/u/1275316631",
	}, {
		Nick: "流浪雪",
		URL:  "http://album.zhenai.com/u/1486293757",
	}, {
		Nick: "不是暖男的暖男",
		URL:  "http://album.zhenai.com/u/1191555599",
	}, {
		Nick: "安然",
		URL:  "http://album.zhenai.com/u/1572218980",
	}, {
		Nick: "簡單是幸福",
		URL:  "http://album.zhenai.com/u/1601974022",
	}, {
		Nick: "寻找有缘人",
		URL:  "http://album.zhenai.com/u/1965268033",
	}, {
		Nick: "艾菲尔的铁塔梦",
		URL:  "http://album.zhenai.com/u/1301136332",
	}, {
		Nick: "墨洁",
		URL:  "http://album.zhenai.com/u/1384450731",
	},
}

const testProfileUrl = "http://album.zhenai.com/u/1584819496"
const testProfileNick = "用户名"

var profileRegExpTestCases = []struct {
	htmlTagString   string
	exceptedProfile profileResult
	description     string
}{
	{
		htmlTagString: fmt.Sprintf(`<a href="%s">%s</a>`,
			testProfileUrl, testProfileNick),
		exceptedProfile: profileResult{
			URL:  testProfileUrl,
			Nick: testProfileNick,
		},
		description: "valid html tag",
	}, {
		htmlTagString: fmt.Sprintf(`<a href="%s" >%s</a>`,
			testProfileUrl, testProfileNick),
		exceptedProfile: profileResult{
			URL:  testProfileUrl,
			Nick: testProfileNick,
		},
		description: "valid html tag with an space tail",
	}, {
		htmlTagString: fmt.Sprintf(`<a target="_blank" href="%s">%s</a>`,
			testProfileUrl, testProfileNick),
		exceptedProfile: profileResult{
			URL:  testProfileUrl,
			Nick: testProfileNick,
		},
		description: "valid html tag with another attribute at head",
	}, {
		htmlTagString: fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`,
			testProfileUrl, testProfileNick),
		exceptedProfile: profileResult{
			URL:  testProfileUrl,
			Nick: testProfileNick,
		},
		description: "valid html tag with another attribute at tail",
	}, {
		htmlTagString: fmt.Sprintf(`<a data-mark="vip" href="%s" target="_blank">%s</a>`,
			testProfileUrl, testProfileNick),
		exceptedProfile: profileResult{
			URL:  testProfileUrl,
			Nick: testProfileNick,
		},
		description: "valid html tag with other attributes at both head and tail",
	}, {
		htmlTagString: fmt.Sprintf(`<a data-mark="vip" href="%s" target="_blank">%s%s%s</a>`,
			testProfileUrl, "\n", testProfileNick, "\n"),
		exceptedProfile: profileResult{
			URL:  testProfileUrl,
			Nick: testProfileNick,
		},
		description: "valid html tag with other attributes in head and tail",
	}, {
		htmlTagString: fmt.Sprintf(`<a data-mark="vip" href="%s" target="_blank">%s%s%s</a>`,
			testProfileUrl, "\n", "用户 名", "\n"),
		exceptedProfile: profileResult{
			URL:  testProfileUrl,
			Nick: "用户 名",
		},
		description: "valid html tag within space",
	},
}
