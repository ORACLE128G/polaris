package model

// anjuke.com profile
type AnjukeProfile struct {
	Id         string
	Name       string
	Price      string
	BuildYears string
	HouseModel string
	Address    string
	Tags       string
	Images     []string
	Entry      interface{}
}
