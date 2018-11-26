package parser

import (
	"polaris/crawler-distributed/config"
	"polaris/crawler/engine/anjuke"
	"polaris/crawler/model"
	"regexp"
)

var nameReg = regexp.MustCompile(`<h2>([^<]+)</h2>`)
var priceReg = regexp.MustCompile(`<em class="sp-price other-price">([^<]+)</em>`)
var buildYearsReg = regexp.MustCompile(`<em class="sp-price other-price">([^<]+)</em>`)
var idReg = regexp.MustCompile(`data-loupan_id="([^"]+)"`)

const types = "anjuke"

func parseProfile(contents []byte, url string) anjuke.ParseResult {
	profile := model.AnjukeProfile{}

	profile.Id = extractString(contents, idReg)
	profile.Name = extractString(contents, nameReg)
	profile.Price = extractString(contents, priceReg)
	profile.BuildYears = extractString(contents, buildYearsReg)

	result := anjuke.ParseResult{
		Items: []anjuke.Item{
			{
				Url:     url,
				Type:    types,
				Id:      profile.Id,
				PayLoad: profile,
			},
		},
	}
	return result
}

func extractString(b []byte, r *regexp.Regexp) string {
	match := r.FindSubmatch(b)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

type ProfileParser struct {
}

// Serialize parser func
func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ProfileParser, nil
}

// Parse profile
func (p *ProfileParser) Parse(contents []byte, url string) anjuke.ParseResult {

	return parseProfile(contents, url)
}

func NewProfileParser() *ProfileParser {
	return &ProfileParser{}

}
