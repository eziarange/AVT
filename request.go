package main

type InitRequest struct {
	Result Result `json:"result"`
	Status string `json:"status"`
}

type Result struct {
	Blocks       Blocks `json:"blocks"`
	IsAuthorized bool   `json:"isAuthorized"`
}

type Blocks struct {
	PersonalImpact PersonalImpact `json:"personalImpact"`
}

type PersonalImpact struct {
	Data Data `json:"data"`
}

type Data struct {
	Co2       interface{} `json:"co2"`
	Energy    interface{} `json:"energy"`
	Water     interface{} `json:"water"`
	Materials int64       `json:"materials"`
	PineYears int64       `json:"pineYears"`
}
