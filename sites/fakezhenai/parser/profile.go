package parser

import (
	"github.com/lihs-learning/go-crawler/engine"
	"github.com/lihs-learning/go-crawler/model"
	"regexp"
	"strconv"
)

var profileAgeRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>([^<]+)岁</td>`)
var profileHeightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>([^<]+)[Cc][Mm]</td>`)
var profileWeightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">([^<]+)[Kk][Gg]</span></td>`)
var profileIncomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var profileMarriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var profileEducationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var profileGenderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
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
}

func ParseProfile(utf8Content []byte, extraInfo ProfileParserExtraInfo) (result engine.ParseResult) {
	profile := model.Profile{
		Name:   extraInfo.Name,
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
	profile.Gender = extractString(utf8Content, profileGenderRe)
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
