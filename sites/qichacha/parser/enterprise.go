package parser

import (
	"regexp"

	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/model"
)

// 基本信息正则
var mainPeopleRe = regexp.MustCompile(
	`<td[^>]*>\s*(?:法定代表人|经营者)\s*</td>[\s\S]*<a .*?href="(/pl_[pra-f0-9]{32}\.html)" class="bname"><h2 class="seo font-20">([^<]*)</h2></a>`)
var registeredCapitalRe = regexp.MustCompile(
	`<td[^>]*>\s*注册资本\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var subscribedCapitalRe = regexp.MustCompile(
	`<td[^>]*>\s*实缴资本\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var operatingStatusRe = regexp.MustCompile(
	`<td[^>]*>\s*经营状态\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var establishedDateRe = regexp.MustCompile(
	`<td[^>]*>\s*成立日期\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var unifiedSocialCreditCodeRe = regexp.MustCompile(
	`<td[^>]*>\s*统一社会信用代码\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var taxpayerIdentificationCodeRe = regexp.MustCompile(
	`<td[^>]*>\s*纳税人识别号\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var registrationCodeRe = regexp.MustCompile(
	`<td[^>]*>\s*注册号\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var organizingInstitutionBarCodeRe = regexp.MustCompile(
	`<td[^>]*>\s*组织机构代码\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var formRe = regexp.MustCompile(
	`<td[^>]*>\s*企业类型\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var industriesClassificationRe = regexp.MustCompile(
	`<td[^>]*>\s*所属行业\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var approvalDateRe = regexp.MustCompile(
	`<td[^>]*>\s*核准日期\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var registrarRe = regexp.MustCompile(
	`<td[^>]*>\s*登记机关\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var locationRe = regexp.MustCompile(
	`<td[^>]*>\s*所属(?:地区|区域)\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var englishNameRe = regexp.MustCompile(
	`<td[^>]*>\s*英文名\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var formerNameRe = regexp.MustCompile(
	`<td[^>]*>\s*曾用名\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var insuredNumberRe = regexp.MustCompile(
	`<td[^>]*>\s*参保人数\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var staffSizeRe = regexp.MustCompile(
	`<td[^>]*>\s*人员规模\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var operatingTermRe = regexp.MustCompile(
	`<td[^>]*>\s*营业期限\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)
var addressRe = regexp.MustCompile(
	`<td[^>]*>\s*企业地址\s*</td>\s*<td[^>]*>\s*(.*?)\s*<`)
var businessScopeRe = regexp.MustCompile(
	`<td[^>]*>\s*经营范围\s*</td>\s*<td[^>]*>\s*(.*?)\s*</td>`)

// 股东信息
// 主要人员

func ParseEnterprise(utf8content []byte, enterpriseBrief EnterpriseBrief) (result engine.ParseResult) {
	enterprise := model.Enterprise{
		ID:   enterpriseBrief.ID,
		Name: enterpriseBrief.Name,
	}

	enterprise.MainPeople = extractString(utf8content, mainPeopleRe)

	match := mainPeopleRe.FindSubmatch(utf8content)
	if len(match) >= 2 {
		// match[1] people page uri
		enterprise.MainPeople = string(match[2])
	}

	enterprise.RegisteredCapital = extractString(utf8content, registeredCapitalRe)
	enterprise.SubscribedCapital = extractString(utf8content, subscribedCapitalRe)
	enterprise.OperatingStatus = extractString(utf8content, operatingStatusRe)
	enterprise.EstablishedDate = extractString(utf8content, establishedDateRe)
	enterprise.UnifiedSocialCreditCode = extractString(utf8content, unifiedSocialCreditCodeRe)
	enterprise.TaxpayerIdentificationCode = extractString(utf8content, taxpayerIdentificationCodeRe)
	enterprise.RegistrationCode = extractString(utf8content, registrationCodeRe)
	enterprise.OrganizingInstitutionBarCode = extractString(utf8content, organizingInstitutionBarCodeRe)
	enterprise.Form = extractString(utf8content, formRe)
	enterprise.IndustriesClassification = extractString(utf8content, industriesClassificationRe)
	enterprise.ApprovalDate = extractString(utf8content, approvalDateRe)
	enterprise.Registrar = extractString(utf8content, registrarRe)
	enterprise.Location = extractString(utf8content, locationRe)
	enterprise.EnglishName = extractString(utf8content, englishNameRe)
	enterprise.FormerName = extractString(utf8content, formerNameRe)
	enterprise.InsuredNumber = extractString(utf8content, insuredNumberRe)
	enterprise.StaffSize = extractString(utf8content, staffSizeRe)
	enterprise.OperatingTerm = extractString(utf8content, operatingTermRe)
	enterprise.Address = extractString(utf8content, addressRe)
	enterprise.BusinessScope = extractString(utf8content, businessScopeRe)

	result.Items = append(result.Items, enterprise)

	return
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
