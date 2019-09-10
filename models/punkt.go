package models

import "encoding/json"

//go:generate msgp

/*
Punkt descibes app collection in MongoDB
*/
type Punkt struct {
	ID            int       `json:"id"`
	City          string    `json:"city"`
	Name          string    `json:"name"`
	Weight        string    `json:"weight"`
	Wholesale     bool      `json:"wholesale"`
	WholesaleText string    `json:"wholesale_text"`
	Mainaddress   string    `json:"mainaddress"`
	Gruppa        int       `json:"gruppa"`
	Address       string    `json:"address"`
	Phone         string    `json:"phone"`
	Phones        []string  `json:"phones"`
	Date          int       `json:"date"`
	Bank          int       `json:"bank"`
	Data          Data      `json:"data"`
	Lat           string    `json:"lat"`
	Lng           string    `json:"lng"`
	Gray          bool      `json:"gray"`
	Actual        bool      `json:"actual"`
	ActualTime    int       `json:"actualTime"`
	Workmodes     Workmodes `json:"workmodes"`
	Workattr      Workattr  `json:"workattr"`
	Map           Map       `json:"map"`
	Mainsort      int       `json:"mainsort"`
}

/*
Data descibes app collection in MongoDB
*/
type Data struct {
	USD  []float64 `json:"USD"`
	EUR  []float64 `json:"EUR"`
	RUB  []float64 `json:"RUB"`
	KGS  []float64 `json:"KGS"`
	CNY  []float64 `json:"CNY"`
	GBP  []float64 `json:"GBP"`
	CHF  []float64 `json:"CHF"`
	UZS  []float64 `json:"UZS"`
	JPY  []float64 `json:"JPY"`
	AUD  []float64 `json:"AUD"`
	TRY  []float64 `json:"TRY"`
	AED  []float64 `json:"AED"`
	UAH  []float64 `json:"UAH"`
	THB  []float64 `json:"THB"`
	INR  []float64 `json:"INR"`
	EGP  []float64 `json:"EGP"`
	CAD  []float64 `json:"CAD"`
	KPW  []float64 `json:"KPW"`
	KRW  []float64 `json:"KRW"`
	MNT  []float64 `json:"MNT"`
	TMT  []float64 `json:"TMT"`
	GEL  []float64 `json:"GEL"`
	GOLD []float64 `json:"GOLD"`
	AZN  []float64 `json:"AZN"`
	BHD  []float64 `json:"BHD"`
	AMD  []float64 `json:"AMD"`
	BYN  []float64 `json:"BYN"`
	BRL  []float64 `json:"BRL"`
	HUF  []float64 `json:"HUF"`
	HKD  []float64 `json:"HKD"`
	DKK  []float64 `json:"DKK"`
	IRR  []float64 `json:"IRR"`
	KWD  []float64 `json:"KWD"`
	MYR  []float64 `json:"MYR"`
	MXN  []float64 `json:"MXN"`
	MDL  []float64 `json:"MDL"`
	NOK  []float64 `json:"NOK"`
	PLN  []float64 `json:"PLN"`
	SAR  []float64 `json:"SAR"`
	XDR  []float64 `json:"XDR"`
	SGD  []float64 `json:"SGD"`
	TJS  []float64 `json:"TJS"`
	CZK  []float64 `json:"CZK"`
	SEK  []float64 `json:"SEK"`
	ZAR  []float64 `json:"ZAR"`
}

/*
Workmodes descibes app collection in MongoDB
*/
type Workmodes struct {
	Mon     []string `json:"mon"`
	Tue     []string `json:"tue"`
	Wed     []string `json:"wed"`
	Thu     []string `json:"thu"`
	Fri     []string `json:"fri"`
	Sat     []string `json:"sat"`
	Sun     []string `json:"sun"`
	Holyday []string `json:"holyday"`
}

/*
Workattr descibes app collection in MongoDB
*/
type Workattr struct {
	Nonstop bool `json:"nonstop"`
	Closed  bool `json:"closed"`
	Worknow bool `json:"worknow"`
}

/*
Map descibes app collection in MongoDB
*/
type Map struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

/*
Marshal Punkt structure to JSON
*/
func (v *Punkt) Marshal() ([]byte, error) {
	res, err := json.Marshal(v)
	return res, err
}

/*
Unmarshal JSON data to Punkt structure
*/
func (v *Punkt) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, v)
	return err
}
