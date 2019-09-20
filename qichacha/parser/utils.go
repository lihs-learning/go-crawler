package parser

import "regexp"

var noneChineseRegexp = regexp.MustCompile(
	"[^\u4e00-\u9fa5]")

func filterChinese(utf8content []byte) []byte {
	return noneChineseRegexp.ReplaceAll(utf8content, []byte{})
}
