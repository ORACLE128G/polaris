package parser

import (
	"polaris/crawler/engine"
	"polaris/crawler/model"
	"regexp"
	"strconv"
)

//var nameExp = regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)
var genderExp = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var ageExp = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightExp = regexp.MustCompile(`<td><span class="label">身高：</span>169CM</td>`)
var weightExp = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
var inComeExp = regexp.MustCompile(`<td><span class="label">月收入：</span>3001-5000元</td>`)
var marriageExp = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationExp = regexp.MustCompile(`<td><span class="label">学历：</span>大学本科</td>`)
var occupationExp = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var huKouExp = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinZuoExp = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseExp = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carExp = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var idExp = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(
	contents [] byte, name string,
	url string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name
	profile.Gender = extractString(contents, genderExp)
	profile.Income = extractString(contents, inComeExp)
	profile.Marriage = extractString(contents, marriageExp)
	profile.Education = extractString(contents, educationExp)
	profile.Occupation = extractString(contents, occupationExp)
	profile.Hukou = extractString(contents, huKouExp)
	profile.Xinzuo = extractString(contents, xinZuoExp)
	profile.House = extractString(contents, houseExp)
	profile.Car = extractString(contents, carExp)
	if age, err := strconv.Atoi(extractString(contents, ageExp)); err == nil {
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractString(contents, heightExp)); err == nil {
		profile.Height = height
	}
	if weight, err := strconv.Atoi(extractString(contents, weightExp)); err == nil {
		profile.Weight = weight
	}
	profile.Marriage = extractString(contents, marriageExp)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idExp),
				PayLoad: profile,
			},
		},
	}
	return result
}

func extractString(contents [] byte, reg *regexp.Regexp) string {
	match := reg.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
