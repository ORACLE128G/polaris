package parser

import (
	"polaris/crawler/engine"
	"polaris/crawler/model"
	"regexp"
	"strconv"
)

var ageRegExp = regexp.MustCompile(
	`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRegExp = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)

func ParseProfile(
	contents [] byte) engine.ParseResult {

	profile := model.Profile{}
	if age, err := strconv.Atoi(extractString(contents, ageRegExp)); err == nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRegExp)
}

func extractString(contents [] byte, reg *regexp.Regexp) string {
	match := reg.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
