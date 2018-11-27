package parser

import (
	"polaris/crawler-distributed/config"
	"polaris/crawler/engine"
	"polaris/crawler/model"
	"regexp"
)

var nameReg = regexp.MustCompile(`<h2>([^<]+)</h2>`)
var priceReg = regexp.MustCompile(`<em class="sp-price other-price">([^<]+)</em>`)
var buildYearsReg = regexp.MustCompile(`<dd><span>([^<]+)</span></dd>`)
var idReg = regexp.MustCompile(`data-loupan_id="([^"]+)"`)
var addressReg = regexp.MustCompile(`<a class="lpAddr-text g-overflow" href="([.]+)" 
soj="loupan_index_map" target="_blank">(^<)+</a>`)

var imgReg = regexp.MustCompile(`<img width="([^"]+)" height="([^"]+)" src="(https://pic4.ajkimg.com/display/[^"]+)">`)

func parseProfile(contents []byte, url string) engine.ParseResult {
	profile := model.AnjukeProfile{}

	profile.Id = extractString(contents, idReg)
	profile.Name = extractString(contents, nameReg)
	profile.Price = extractString(contents, priceReg)
	profile.BuildYears = extractString(contents, buildYearsReg)
	profile.Address = extractString(contents, addressReg)
	profile.Images = extractImages(contents, imgReg)
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:    url,
				// todo: anjuke
				Type:   "zhenai",
				Id:     profile.Id,
				Anjuke: profile,
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

func extractImages(b []byte, reg *regexp.Regexp) []string {

	match := reg.FindAllSubmatch(b, -1)
	var r []string
	for _, e := range match {
		r = append(r, string(e[3]))
	}
	return r
}

type ProfileParser struct {
}

// Serialize parser func
func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ProfileParser, nil
}

// Parse profile
func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {

	return parseProfile(contents, url)
}

func NewProfileParser() *ProfileParser {
	return &ProfileParser{}

}
