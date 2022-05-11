package main

import (
	"math/rand"
	"time"

	vehicle "github.com/CK-Gnap/traffic/vehicles"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	road := vehicle.Road{
		Name: "Mukmontri",
	}

	vehicles := []vehicle.Vehicle{
		&vehicle.Car{Brand: "BMW", Model: "2 Series Gran Coupé", Year: 2019, Wheels: 4, Color: "Black"},
		&vehicle.Truck{Brand: "FOTON", Model: "Auman EST A", Year: 2021, Wheels: 10, Color: "White"},
		&vehicle.Motorcycle{Brand: "Yamaha", Model: "XMAX SP", Year: 2022, Wheels: 2, Color: "Blue"},
		&vehicle.Truck{Brand: "FOTON", Model: "Auman EST M 12T", Year: 2019, Wheels: 6, Color: "Black"},
		&vehicle.Car{Brand: "BMW", Model: "X1", Year: 2018, Wheels: 4, Color: "Brown"},
		&vehicle.Motorcycle{Brand: "Yamaha", Model: "R7", Year: 2021, Wheels: 2, Color: "Pink"},
		&vehicle.Car{Brand: "BMW", Model: "7 Series Sedan", Year: 2021, Wheels: 4, Color: "Purple"},
		&vehicle.Truck{Brand: "FOTON", Model: "AUV", Year: 2020, Wheels: 4, Color: "Yellow"},
		&vehicle.Motorcycle{Brand: "Yamaha", Model: "Finn", Year: 2021, Wheels: 2, Color: "Green"},
		&vehicle.Motorcycle{Brand: "Yamaha", Model: "XSR155", Year: 2021, Wheels: 2, Color: "Black"},
	}

	err := road.SimulateTraffic(vehicles)
	if err != nil {
		panic(err)
	}
}
