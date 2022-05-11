package vehicles

import (
	"testing"
)

func TestGetTrafficLightColor(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{
			name:  "when get green color is happy",
			value: 1,
			want:  "Green",
		},
		{
			name:  "when get yellow color is happy",
			value: 2,
			want:  "Yellow",
		},
		{
			name:  "when get red color is happy",
			value: 3,
			want:  "Red",
		},
		{
			name:  "when get traffic light out of range",
			value: 4,
			want:  "Out of range",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getTrafficLightColor(test.value)
			if got != test.want {
				t.Errorf("getTrafficLightColor() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestImportVehicles(t *testing.T) {
	vehicles := []Vehicle{
		&Car{Brand: "BMW", Model: "2 Series Gran Coupé", Year: 2019, Wheels: 4, Color: "Black"},
		&Truck{Brand: "FOTON", Model: "Auman EST A", Year: 2021, Wheels: 10, Color: "White"},
		&Motorcycle{Brand: "Yamaha", Model: "XMAX SP", Year: 2022, Wheels: 2, Color: "Blue"},
	}

	type args struct {
		trafficLightIn int
		vehicles       []Vehicle
	}

	tests := []struct {
		name    string
		args    args
		isValue bool
		isError bool
	}{
		{
			name: "when happy with green light",
			args: args{
				trafficLightIn: 1,
				vehicles:       vehicles,
			},
			isValue: true,
			isError: false,
		},
		{
			name: "when happy with yellow light",
			args: args{
				trafficLightIn: 2,
				vehicles:       vehicles,
			},
			isValue: true,
			isError: false,
		},
		{
			name: "when happy with red light",
			args: args{
				trafficLightIn: 3,
				vehicles:       vehicles,
			},
			isValue: true,
			isError: false,
		},
		{
			name: "when error Traffic Light In is out of range",
			args: args{
				trafficLightIn: 33,
				vehicles:       vehicles,
			},
			isValue: false,
			isError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			road := Road{
				Name: "Mukmontri",
			}

			got, err := road.ImportVehicles(test.args.trafficLightIn, test.args.vehicles)
			if (err != nil) != test.isError {
				t.Errorf("ImportVehicles() error = %v, isError %v", err, test.isError)
			} else {
				if !(got < 0) != test.isValue {
					t.Errorf("ImportVehicles() = %v, isValue %v", got, test.isValue)
				}
			}
		})
	}
}

func TestExportVehicles(t *testing.T) {
	vehicles := []Vehicle{
		&Car{Brand: "BMW", Model: "2 Series Gran Coupé", Year: 2019, Wheels: 4, Color: "Black"},
		&Truck{Brand: "FOTON", Model: "Auman EST A", Year: 2021, Wheels: 10, Color: "White"},
		&Motorcycle{Brand: "Yamaha", Model: "XMAX SP", Year: 2022, Wheels: 2, Color: "Blue"},
	}

	type args struct {
		trafficLightOut int
		vehicles        []Vehicle
	}

	tests := []struct {
		name    string
		args    args
		isValue bool
		isError bool
	}{
		{
			name: "when happy with green light",
			args: args{
				trafficLightOut: 1,
				vehicles:        vehicles,
			},
			isValue: true,
			isError: false,
		},
		{
			name: "when happy with yellow light",
			args: args{
				trafficLightOut: 2,
				vehicles:        vehicles,
			},
			isValue: true,
			isError: false,
		},
		{
			name: "when happy with red light",
			args: args{
				trafficLightOut: 3,
				vehicles:        vehicles,
			},
			isValue: true,
			isError: false,
		},
		{
			name: "when error Traffic Light Out is out of range",
			args: args{
				trafficLightOut: 4,
				vehicles:        vehicles,
			},
			isValue: false,
			isError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			road := Road{
				Name:     "Mukmontri",
				Vehicles: test.args.vehicles,
			}

			got, err := road.ExportVehicles(test.args.trafficLightOut)
			if (err != nil) != test.isError {
				t.Errorf("ExportVehicles() error = %v, isError %v", err, test.isError)
			} else {
				if !(got < 0) != test.isValue {
					t.Errorf("ExportVehicles() = %v, isValue %v", got, test.isValue)
				}
			}
		})
	}
}
