package parser

const ProvincesSize = 31

var provinceTests = []struct {
	Name string
	URL  string
}{
	{
		Name: "安徽",
		URL:  "https://www.qichacha.com/g_AH.html",
	},
	{
		Name: "北京",
		URL:  "https://www.qichacha.com/g_BJ.html",
	},
	{
		Name: "重庆",
		URL:  "https://www.qichacha.com/g_CQ.html",
	},
	{
		Name: "福建",
		URL:  "https://www.qichacha.com/g_FJ.html",
	},
	{
		Name: "甘肃",
		URL:  "https://www.qichacha.com/g_GS.html",
	},
	{
		Name: "广东",
		URL:  "https://www.qichacha.com/g_GD.html",
	},
	{
		Name: "广西",
		URL:  "https://www.qichacha.com/g_GX.html",
	},
	{
		Name: "贵州",
		URL:  "https://www.qichacha.com/g_GZ.html",
	},
	{
		Name: "海南",
		URL:  "https://www.qichacha.com/g_HAIN.html",
	},
	{
		Name: "河北",
		URL:  "https://www.qichacha.com/g_HB.html",
	},
	{
		Name: "黑龙江",
		URL:  "https://www.qichacha.com/g_HLJ.html",
	},
	{
		Name: "河南",
		URL:  "https://www.qichacha.com/g_HEN.html",
	},
	{
		Name: "湖北",
		URL:  "https://www.qichacha.com/g_HUB.html",
	},
	{
		Name: "湖南",
		URL:  "https://www.qichacha.com/g_HUN.html",
	},
	{
		Name: "江苏",
		URL:  "https://www.qichacha.com/g_JS.html",
	},
	{
		Name: "江西",
		URL:  "https://www.qichacha.com/g_JX.html",
	},
	{
		Name: "吉林",
		URL:  "https://www.qichacha.com/g_JL.html",
	},
	{
		Name: "辽宁",
		URL:  "https://www.qichacha.com/g_LN.html",
	},
	{
		Name: "内蒙古",
		URL:  "https://www.qichacha.com/g_NMG.html",
	},
	{
		Name: "宁夏",
		URL:  "https://www.qichacha.com/g_NX.html",
	},
	{
		Name: "青海",
		URL:  "https://www.qichacha.com/g_QH.html",
	},
	{
		Name: "山东",
		URL:  "https://www.qichacha.com/g_SD.html",
	},
	{
		Name: "上海",
		URL:  "https://www.qichacha.com/g_SH.html",
	},
	{
		Name: "山西",
		URL:  "https://www.qichacha.com/g_SX.html",
	},
	{
		Name: "陕西",
		URL:  "https://www.qichacha.com/g_SAX.html",
	},
	{
		Name: "四川",
		URL:  "https://www.qichacha.com/g_SC.html",
	},
	{
		Name: "天津",
		URL:  "https://www.qichacha.com/g_TJ.html",
	},
	{
		Name: "新疆",
		URL:  "https://www.qichacha.com/g_XJ.html",
	},
	{
		Name: "西藏",
		URL:  "https://www.qichacha.com/g_XZ.html",
	},
	{
		Name: "云南",
		URL:  "https://www.qichacha.com/g_YN.html",
	},
	{
		Name: "浙江",
		URL:  "https://www.qichacha.com/g_ZJ.html",
	},
}
