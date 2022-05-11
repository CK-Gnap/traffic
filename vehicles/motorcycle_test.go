package vehicles

import (
	"testing"
)

func TestStartMotorcycle(t *testing.T) {
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
			motorcycle := Motorcycle{}
			got := motorcycle.Start()
			if (got > 0) != test.isValue {
				t.Errorf("motorcycle.Start() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestDriveMotorcycle(t *testing.T) {
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
			motorcycle := Motorcycle{}
			_ = motorcycle.Start()
			got := motorcycle.Drive()
			if (got > 0) != test.isValue {
				t.Errorf("motorcycle.Drive() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestSpeedUpMotorcycle(t *testing.T) {
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
			motorcycle := Motorcycle{}
			_ = motorcycle.Start()
			_ = motorcycle.Drive()
			got := motorcycle.SpeedUp()
			if (got > 0) != test.isValue {
				t.Errorf("motorcycle.SpeedUp() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestSpeedDownMotorcycle(t *testing.T) {
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
			motorcycle := Motorcycle{}
			_ = motorcycle.Start()
			_ = motorcycle.Drive()
			got := motorcycle.SpeedDown()
			if (got > 0) != test.isValue {
				t.Errorf("motorcycle.SpeedDown() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}

func TestStopMotorcycle(t *testing.T) {
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
			motorcycle := Motorcycle{}
			_ = motorcycle.Start()
			_ = motorcycle.Drive()
			_ = motorcycle.SpeedUp()
			_ = motorcycle.SpeedDown()
			got := motorcycle.Stop()
			if (got != 0) != test.isValue {
				t.Errorf("motorcycle.Stop() = %v, isValue %v", got, test.isValue)
			}
		})
	}
}
