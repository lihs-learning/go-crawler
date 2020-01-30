package parser

import (
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/model"
	"regexp"
	"strconv"
)

var profileAgeRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>(\d+)岁</div>`)
var profileHeightRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>(\d+)[cC][mM]</div>`)
var profileWeightRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>(\d+)[kK][gG]</div>`)
var profileIncomeRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>月收入[:：]([^<]+)</div>`)
var profileMarriageRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>(未婚|离异|丧偶)</div>`)
var profileEducationRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>(高中及以下|中专|大专|大学本科|硕士|博士)</div>`)
var profileOccupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var profileResidenceRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var profileConstellationRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var profileHouseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var profileCarRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

//var profileGuessRe = regexp.MustCompile(
//	`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
//var profileIdUrlRe = regexp.MustCompile(
//	`http://album.zhenai.com/u/([\d]+)`)

type ProfileParserExtraInfo struct {
	Name   string
	Gender string
}

func ParseProfile(utf8Content []byte, extraInfo ProfileParserExtraInfo) (result engine.ParseResult) {
	profile := model.Profile{
		Name:   extraInfo.Name,
		Gender: extraInfo.Gender,
	}

	age, err := strconv.Atoi(
		extractString(utf8Content, profileAgeRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(utf8Content, profileHeightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(
		extractString(utf8Content, profileWeightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(utf8Content, profileIncomeRe)
	profile.Marriage = extractString(utf8Content, profileMarriageRe)
	profile.Education = extractString(utf8Content, profileEducationRe)
	profile.Occupation = extractString(utf8Content, profileOccupationRe)
	profile.Residence = extractString(utf8Content, profileResidenceRe)
	profile.Constellation = extractString(utf8Content, profileConstellationRe)
	profile.House = extractString(utf8Content, profileHouseRe)
	profile.Car = extractString(utf8Content, profileCarRe)

	result.Items = []interface{}{
		profile,
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
