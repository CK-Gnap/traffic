package vehicles

import (
	"testing"
)

func TestStartTruck(t *testing.T) {
	tests := []struct {
		name    string
		isValue bool
	}{
		{
			name:    "when happy",
			isValue: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			truck := Truck{}
			got := truck.Start()
			if (got > 0) != test.isValue {
				t.Errorf("truck.Start() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestDriveTruck(t *testing.T) {
	tests := []struct {
		name    string
		isValue bool
	}{
		{
			name:    "when happy",
			isValue: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			truck := Truck{}
			_ = truck.Start()
			got := truck.Drive()
			if (got > 0) != test.isValue {
				t.Errorf("truck.Drive() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestSpeedUpTruck(t *testing.T) {
	tests := []struct {
		name    string
		isValue bool
	}{
		{
			name:    "when happy",
			isValue: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			truck := Truck{}
			_ = truck.Start()
			_ = truck.Drive()
			got := truck.SpeedUp()
			if (got > 0) != test.isValue {
				t.Errorf("truck.SpeedUp() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestSpeedDownTruck(t *testing.T) {
	tests := []struct {
		name    string
		isValue bool
	}{
		{
			name:    "when happy",
			isValue: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			truck := Truck{}
			_ = truck.Start()
			_ = truck.Drive()
			got := truck.SpeedDown()
			if (got > 0) != test.isValue {
				t.Errorf("truck.SpeedDown() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestStopTruck(t *testing.T) {
	tests := []struct {
		name    string
		isValue bool
	}{
		{
			name:    "when happy",
			isValue: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			truck := Truck{}
			_ = truck.Start()
			_ = truck.Drive()
			_ = truck.SpeedUp()
			_ = truck.SpeedDown()
			got := truck.Stop()
			if (got != 0) != test.isValue {
				t.Errorf("truck.Stop() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}
