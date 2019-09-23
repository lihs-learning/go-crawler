package parser

import (
	"github.com/lihs-learning/go-crawler/engine"
	"regexp"
)

// 基本信息正则
var mainPeopleRe = regexp.MustCompile(
	`<td[^>]*>\s*(?:法定代表人|经营者)\s*</td>[\s\S]*<a .*?href="(/pl_[pa-f0-9]{32}\.html)" class="bname"><h2 class="seo font-20">([^<]*)</h2></a>`)
var registeredCapital = regexp.MustCompile(
	`<td[^>]*>\s*注册资本\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var subscribedCapital = regexp.MustCompile(
	`<td[^>]*>\s*实缴资本\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var operatingStatus = regexp.MustCompile(
	`<td[^>]*>\s*经营状态\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var establishedDate = regexp.MustCompile(
	`<td[^>]*>\s*成立日期\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var unifiedSocialCreditCode = regexp.MustCompile(
	`<td[^>]*>\s*统一社会信用代码\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var taxpayerIdentificationCode = regexp.MustCompile(
	`<td[^>]*>\s*纳税人识别号\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var registrationCode = regexp.MustCompile(
	`<td[^>]*>\s*注册号\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var organizingInstitutionBarCode = regexp.MustCompile(
	`<td[^>]*>\s*组织机构代码\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var form = regexp.MustCompile(
	`<td[^>]*>\s*企业类型\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var industriesClassification = regexp.MustCompile(
	`<td[^>]*>\s*所属行业\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var approvalDate = regexp.MustCompile(
	`<td[^>]*>\s*核准日期\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var registrar = regexp.MustCompile(
	`<td[^>]*>\s*登记机关\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var location = regexp.MustCompile(
	`<td[^>]*>\s*所属区域\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var englishName = regexp.MustCompile(
	`<td[^>]*>\s*英文名\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var formerName = regexp.MustCompile(
	`<td[^>]*>\s*曾用名\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var insuredNumber = regexp.MustCompile(
	`<td[^>]*>\s*参保人数\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var staffSize = regexp.MustCompile(
	`<td[^>]*>\s*人员规模\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var operatingTerm = regexp.MustCompile(
	`<td[^>]*>\s*营业期限\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var address = regexp.MustCompile(
	`<td[^>]*>\s*企业地址\s*</td>\s*<td[^>]*>\s*(.*?)\s*<`)
var businessScope = regexp.MustCompile(
	`<td[^>]*>\s*经营范围\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)

// 股东信息
// 主要人员

func ParseEnterprise(utf8content []byte) (result engine.ParseResult) {

	return
}
