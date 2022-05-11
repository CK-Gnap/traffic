package vehicles

import (
	"testing"
)

func TestStartCar(t *testing.T) {
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
			car := Car{}
			got := car.Start()
			if (got > 0) != test.isValue {
				t.Errorf("car.Start() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestDriveCar(t *testing.T) {
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
			car := Car{}
			_ = car.Start()
			got := car.Drive()
			if (got > 0) != test.isValue {
				t.Errorf("car.Drive() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestSpeedUpCar(t *testing.T) {
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
			car := Car{}
			_ = car.Start()
			_ = car.Drive()
			got := car.SpeedUp()
			if (got > 0) != test.isValue {
				t.Errorf("car.SpeedUp() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestSpeedDownCar(t *testing.T) {
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
			car := Car{}
			_ = car.Start()
			_ = car.Drive()
			got := car.SpeedDown()
			if (got > 0) != test.isValue {
				t.Errorf("car.SpeedDown() = %v, isValue %v", got > 0, test.isValue)
			}
		})
	}
}

func TestStopCar(t *testing.T) {
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
			car := Car{}
			_ = car.Start()
			_ = car.Drive()
			_ = car.SpeedUp()
			_ = car.SpeedDown()
			got := car.Stop()
			if (got != 0) != test.isValue {
				t.Errorf("car.Stop() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}
