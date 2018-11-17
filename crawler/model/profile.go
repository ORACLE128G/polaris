package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hukou      string
	Xinzuo     string
	House      string
	Car        string
	Entry      []string
}

func FromJson(i interface{}) (Profile, error) {
	var profile Profile
	r, err := json.Marshal(i)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(r, &profile)
	return profile, err
}
