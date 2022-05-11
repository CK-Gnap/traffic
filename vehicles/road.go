package vehicles

import (
	"log"
)

const (
	GREEN  = 1
	YELLOW = 2
	RED    = 3
)

type Road struct {
	Name     string
	Vehicles []Vehicle
}

type TrafficLight struct{}

func (road *Road) VehiclesIn(vehicles []Vehicle) int {
	log.Printf(">>>>>> Vehicles Start to %s Road <<<<<<\n", road.Name)
	len := 0
	for _, vehicle := range vehicles {
		_ = vehicle.Start()
		_ = vehicle.Drive()
		road.Vehicles = append(road.Vehicles, vehicle)
		len++
	}
	return len
}

func (road *Road) VehiclesInWaitingToStop(vehicles []Vehicle) {
	log.Printf(">>>>>> Vehicles Speed Down on Other Road <<<<<<\n")
	log.Printf(">>>>>>>>> Vehicles Stop on Other Road <<<<<<<<<\n")
	for _, vehicle := range vehicles {
		_ = vehicle.Start()
		_ = vehicle.SpeedUp()
		_ = vehicle.Drive()
		_ = vehicle.SpeedDown()
		_ = vehicle.Stop()
	}
}
func (road *Road) VehiclesInStop(vehicles []Vehicle) {
	log.Printf(">>>>>>>>> Vehicles Stop on Other Road <<<<<<<<<\n")
	for _, vehicle := range vehicles {
		_ = vehicle.Start()
		_ = vehicle.Drive()
		_ = vehicle.SpeedDown()
		_ = vehicle.Stop()
	}
}

func (road *Road) VehiclesOut() int {
	log.Printf("<<<<<< Vehicles Exit from %s Road >>>>>>\n", road.Name)
	len := 0
	for index, vehicle := range road.Vehicles {
		_ = vehicle.Start()
		_ = vehicle.Drive()
		_ = vehicle.SpeedUp()
		road.RemoveVehicle(index)
		len++
	}
	return len
}

func (road *Road) VehiclesOutWaitingToStop() {
	log.Printf(">>>>>> Vehicles Speed Down on %s Road <<<<<<\n", road.Name)
	log.Printf(">>>>>> Vehicles Stop on %s Road <<<<<<\n", road.Name)
	for index, vehicle := range road.Vehicles {
		_ = vehicle.Start()
		_ = vehicle.Drive()
		_ = vehicle.SpeedDown()
		_ = vehicle.Stop()
		road.RemoveVehicle(index)
	}
}
func (road *Road) VehiclesOutStop() {
	log.Printf(">>>>>> Vehicles Stop %s Road <<<<<<\n", road.Name)
	for index, vehicle := range road.Vehicles {
		_ = vehicle.Start()
		_ = vehicle.SpeedDown()
		_ = vehicle.Stop()
		road.RemoveVehicle(index)
	}
}

func (road *Road) RemoveVehicle(index int) {
	vehicles := road.Vehicles
	if len(vehicles) > 0 {
		road.Vehicles[index] = vehicles[len(vehicles)-1]
	}
}

func (t TrafficLight) RandomTrafficLight(value int) int {
	var trafficLight int
	if value == 1 {
		trafficLight = GREEN
	} else if value == 2 {
		trafficLight = YELLOW
	} else if value == 3 {
		trafficLight = RED
	} else {
		trafficLight = 0
	}
	return trafficLight
}
