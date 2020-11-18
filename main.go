package main

import (
	"fmt"
	"github.com/SyedMuzamiM/software_simple/data"
)

func main() {
	fmt.Println("Welcome to the app")
	grower := data.Grower{
		ID: "F0123",
		Name: "Nazir Ahmad",
		Adress: "Sherpora",
		Phone: 8493886280,
		Delivery: data.Delivery{
				ID: 1,
				DriverName: "Syed",
				VechileNo: "JK13F 0482",
				Mobile: 8082008463,
				ItemQty: 120,
				ItemName: "Apple",
			},
	}
	fmt.Println(grower.ID, grower.Deliveries.ID)
}