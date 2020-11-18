package data

import (
	"encoding/json"
	"io"
)

// Grower defines the structure for an API Grower
type Grower struct {
	ID       string
	Name     string
	Adress   string
	Phone    int64
	Delivery Delivery
}

// Delivery definers the structure for the refrence of the Grower
type Delivery struct {
	ID         int
	DriverName string
	Mobile     int
	VechileNo  string
	ItemQty    int
	ItemName   string
	Date       string
}

// FromJSON forms the JSON of the Grower
func (g *Grower) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(g)
}

// Growers is the collection of Grower
type Growers []*Grower

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (g *Growers) ToJSON(r io.Writer) error {
	e := json.NewEncoder(r)
	return e.Encode(g)
}

// GetGrowers returns a list of Growers
func GetGrowers() Growers {
	return growerList
}

var growerList = []*Grower{
	&Grower{
		ID:     "F0123",
		Name:   "Nazir Ahmad",
		Adress: "Sherpora",
		Phone:  8493886280,
		Delivery: Delivery{
			ID:         1,
			DriverName: "Syed",
			VechileNo:  "JK13F 0482",
			Mobile:     8082008463,
			ItemQty:    120,
			ItemName:   "Apple",
		},
	},
	&Grower{
		ID:     "F0133",
		Name:   "Fayaz Dar",
		Adress: "Akhal",
		Phone:  8493886280,
		Delivery: Delivery{
			ID:         1,
			DriverName: "Imran",
			VechileNo:  "JK13F 0082",
			Mobile:     8082008463,
			ItemQty:    100,
			ItemName:   "Apple",
		},
	},
}
