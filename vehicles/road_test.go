package vehicles

import "testing"

func TestVehiclesIn(t *testing.T) {
	vehicles := []Vehicle{
		&Car{Brand: "BMW", Model: "2 Series Gran Coupé", Year: 2019, Wheels: 4, Color: "Black"},
		&Truck{Brand: "FOTON", Model: "Auman EST A", Year: 2021, Wheels: 10, Color: "White"},
		&Motorcycle{Brand: "Yamaha", Model: "XMAX SP", Year: 2022, Wheels: 2, Color: "Blue"},
	}

	tests := []struct {
		name     string
		vehicles []Vehicle
		want     int
	}{
		{
			name:     "when happy",
			vehicles: vehicles,
			want:     3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			road := Road{
				Name: "Mukmontri",
			}
			got := road.VehiclesIn(test.vehicles)
			if got != test.want {
				t.Errorf("VehiclesIn() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestVehiclesInWaitingToStop(t *testing.T) {
	vehicles := []Vehicle{
		&Car{Brand: "BMW", Model: "2 Series Gran Coupé", Year: 2019, Wheels: 4, Color: "Black"},
		&Truck{Brand: "FOTON", Model: "Auman EST A", Year: 2021, Wheels: 10, Color: "White"},
		&Motorcycle{Brand: "Yamaha", Model: "XMAX SP", Year: 2022, Wheels: 2, Color: "Blue"},
	}

	tests := []struct {
		name     string
		vehicles []Vehicle
	}{
		{
			name:     "when happy",
			vehicles: vehicles,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			road := Road{
				Name: "Mukmontri",
			}
			road.VehiclesInWaitingToStop(test.vehicles)
		})
	}
}

func TestVehiclesInStop(t *testing.T) {
	vehicles := []Vehicle{
		&Car{Brand: "BMW", Model: "2 Series Gran Coupé", Year: 2019, Wheels: 4, Color: "Black"},
		&Truck{Brand: "FOTON", Model: "Auman EST A", Year: 2021, Wheels: 10, Color: "White"},
		&Motorcycle{Brand: "Yamaha", Model: "XMAX SP", Year: 2022, Wheels: 2, Color: "Blue"},
	}

	tests := []struct {
		name     string
		vehicles []Vehicle
	}{
		{
			name:     "when happy",
			vehicles: vehicles,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			road := Road{
				Name: "Mukmontri",
			}
			road.VehiclesInStop(test.vehicles)
		})
	}
}

func TestVehiclesOut(t *testing.T) {
	vehicles := []Vehicle{
		&Car{Brand: "BMW", Year: 2019, Wheels: 4, Color: "Black"},
		&Truck{Brand: "FOTON", Year: 2019, Wheels: 4, Color: "Black"},
		&Motorcycle{Brand: "Yamaha", Year: 2019, Wheels: 2, Color: "Black"},
	}

	tests := []struct {
		name     string
		vehicles []Vehicle
		want     int
	}{
		{
			name:     "when happy",
			vehicles: vehicles,
			want:     3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			road := Road{
				Name:     "Mukmontri",
				Vehicles: test.vehicles,
			}
			got := road.VehiclesOut()
			if got != test.want {
				t.Errorf("VehiclesOut() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestVehiclesOutWaitingToStop(t *testing.T) {
	vehicles := []Vehicle{
		&Car{Brand: "BMW", Model: "2 Series Gran Coupé", Year: 2019, Wheels: 4, Color: "Black"},
		&Truck{Brand: "FOTON", Model: "Auman EST A", Year: 2021, Wheels: 10, Color: "White"},
		&Motorcycle{Brand: "Yamaha", Model: "XMAX SP", Year: 2022, Wheels: 2, Color: "Blue"},
	}

	tests := []struct {
		name     string
		vehicles []Vehicle
	}{
		{
			name:     "when happy waiting stop",
			vehicles: vehicles,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			road := Road{
				Name:     "Mukmontri",
				Vehicles: test.vehicles,
			}
			road.VehiclesOutWaitingToStop()
		})
	}
}

func TestVehiclesOutStop(t *testing.T) {
	vehicles := []Vehicle{
		&Car{Brand: "BMW", Model: "2 Series Gran Coupé", Year: 2019, Wheels: 4, Color: "Black"},
		&Truck{Brand: "FOTON", Model: "Auman EST A", Year: 2021, Wheels: 10, Color: "White"},
		&Motorcycle{Brand: "Yamaha", Model: "XMAX SP", Year: 2022, Wheels: 2, Color: "Blue"},
	}

	tests := []struct {
		name     string
		vehicles []Vehicle
	}{
		{
			name:     "when happy stop",
			vehicles: vehicles,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			road := Road{
				Name:     "Mukmontri",
				Vehicles: test.vehicles,
			}
			road.VehiclesOutStop()
		})
	}
}

func TestRemoveVehicle(t *testing.T) {
	vehicles := []Vehicle{
		&Car{Brand: "BMW", Year: 2019, Wheels: 4, Color: "Black"},
		&Truck{Brand: "FOTON", Year: 2019, Wheels: 4, Color: "Black"},
		&Motorcycle{Brand: "Yamaha", Year: 2019, Wheels: 2, Color: "Black"},
	}

	tests := []struct {
		name     string
		vehicles []Vehicle
	}{
		{
			name:     "when happy remove vehicle",
			vehicles: vehicles,
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			road := Road{
				Name: "Mukmontri",
			}
			road.RemoveVehicle(i)
		})
	}
}

func TestRandomTrafficLight(t *testing.T) {
	trafficLight := TrafficLight{}
	tests := []struct {
		name    string
		randInt int
		want    int
	}{
		{
			name:    "when happy Green",
			randInt: 1,
			want:    1,
		},
		{
			name:    "when happy Yellow",
			randInt: 2,
			want:    2,
		},
		{
			name:    "when happy Red",
			randInt: 3,
			want:    3,
		},
		{
			name:    "when happy but out of range",
			randInt: 25,
			want:    0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := trafficLight.RandomTrafficLight(test.randInt)
			if got != test.want {
				t.Errorf("trafficLight.RandomTrafficLight() = %v, want %v", got, test.want)
			}
		})
	}
}
