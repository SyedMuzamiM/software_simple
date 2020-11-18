package data

import (
)

type Grower struct {
	ID string
	Name string
	Adress string
	Phone int64
	Delivery
}

type Delivery struct {
	ID int
	DriverName string
	Mobile int
	VechileNo string
	ItemQty int
	ItemName string
	Date string
}



func main() {
	return
}