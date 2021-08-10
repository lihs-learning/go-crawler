package parser

import "fmt"

type profileResult struct {
	Nick string
	URL  string
}

var cityTestCases = []profileResult{
	{
		Nick: "寂寞成影萌宝",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
	}, {
		Nick: "何必怀念深碍",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/4143921402121684230",
	}, {
		Nick: "猖狂小丸子",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/8020950579916830882",
	}, {
		Nick: "归戾柠小萌",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/3119949895472202650",
	}, {
		Nick: "逍遥浪子晴妹儿.",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/93560234177530413",
	}, {
		Nick: "今朝醉者蓝莓布朗",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/4964470582511013139",
	}, {
		Nick: "简白草莓裙摆",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/3765593846236701995",
	}, {
		Nick: "独久厌闹晴妹儿.",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/1991290594400503513",
	}, {
		Nick: "野性入骨蓝莓布朗",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/9002337925826124089",
	}, {
		Nick: "一见不钟情baby",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/7277509863024900989",
	}, {
		Nick: "称霸全服病娇",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/3211812056555942094",
	}, {
		Nick: "寂寞成影迁就",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/52340934278003036",
	}, {
		Nick: "海阔天空遇到",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/4254843632044688453",
	}, {
		Nick: "原来无话可说胆小鬼",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/7269920654214480798",
	}, {
		Nick: "无戏配角爱你",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/5366828293671555101",
	}, {
		Nick: "高冷绅士baby",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/6112248673500607439",
	}, {
		Nick: "一身傲气小丸子",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/8723163754173477951",
	}, {
		Nick: "断念欣欣雨",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/3353546636144107915",
	}, {
		Nick: "逍遥べ无痕莓哒",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/2858527851503377365",
	}, {
		Nick: "限时拥抱万能萌妹",
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/5785614884937613464",
	},
}

const testProfileUrl = "http://localhost:8080/mock/album.zhenai.com/u/1584819496"
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
