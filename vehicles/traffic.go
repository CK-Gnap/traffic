package vehicles

import (
	"fmt"
	"log"

	helper "github.com/CK-Gnap/traffic/helpers"
)

func (road *Road) SimulateTraffic(vehicles []Vehicle) error {

	trafficLight := TrafficLight{}
	vehiclesIn := 0
	errVehiclesIn := error(nil)
	vehiclesOut := 0
	errVehiclesOut := error(nil)
	vehiclesOnRoad := 0
	lineMsg := "==============================================="

	log.Println(lineMsg)
	trafficLightIn := trafficLight.RandomTrafficLight(helper.RandInt(1, 3))
	trafficLightInColor := getTrafficLightColor(trafficLightIn)
	log.Printf("Traffic Light In is : %s\n", trafficLightInColor)
	log.Println(lineMsg)
	if len(road.Vehicles) == 0 {
		vehiclesIn, errVehiclesIn = road.ImportVehicles(trafficLightIn, vehicles)
		if errVehiclesIn != nil {
			return errVehiclesIn
		}

	}

	log.Println(lineMsg)
	trafficLightOut := trafficLight.RandomTrafficLight(helper.RandInt(1, 3))
	trafficLightOutColor := getTrafficLightColor(trafficLightOut)
	log.Printf("Traffic Light Out is : %s\n", trafficLightOutColor)
	log.Println(lineMsg)
	if len(road.Vehicles) > 0 {
		vehiclesOut, errVehiclesOut = road.ExportVehicles(trafficLightOut)
		if errVehiclesOut != nil {
			return errVehiclesOut
		}
	}

	vehiclesOnRoad = helper.Minus(vehiclesIn, vehiclesOut)
	log.Println(lineMsg)
	log.Printf("Vehicles Total : %v\n", len(vehicles))
	log.Printf("Vehicles In : %v\n", vehiclesIn)
	log.Printf("Vehicles Out : %v\n", vehiclesOut)
	log.Printf("Vehicles on %s Road : %v\n", road.Name, vehiclesOnRoad)
	log.Println(lineMsg)
	return nil
}

func (road *Road) ImportVehicles(trafficLightIn int, vehicles []Vehicle) (int, error) {
	vehiclesIn := 0
	switch trafficLightIn {
	case GREEN:
		randomVehiclesIn := helper.RandInt(1, len(vehicles))
		vehiclesIn = road.VehiclesIn(vehicles[:randomVehiclesIn])
		if len(vehicles) != randomVehiclesIn {
			road.VehiclesInWaitingToStop(vehicles[randomVehiclesIn:])
		}
	case YELLOW:
		road.VehiclesInWaitingToStop(vehicles)
	case RED:
		road.VehiclesInStop(vehicles)
	default:
		return -1, fmt.Errorf("Traffic Light In is out of range")
	}

	return vehiclesIn, nil
}

func (road *Road) ExportVehicles(trafficLightOut int) (int, error) {
	vehiclesOut := 0
	switch trafficLightOut {
	case GREEN:
		vehicles := road.Vehicles
		randomVehiclesOut := helper.RandInt(1, len(vehicles))
		road.Vehicles = vehicles[:randomVehiclesOut]
		vehiclesOut = road.VehiclesOut()
		if len(vehicles) != randomVehiclesOut {
			road.Vehicles = vehicles[randomVehiclesOut:]
			road.VehiclesOutWaitingToStop()
		}
	case YELLOW:
		road.VehiclesOutWaitingToStop()
	case RED:
		road.VehiclesOutStop()
	default:
		return -1, fmt.Errorf("Traffic Light Out is out of range")
	}

	return vehiclesOut, nil
}

func getTrafficLightColor(value int) string {
	var color string
	if value == 1 {
		color = "Green"
	} else if value == 2 {
		color = "Yellow"
	} else if value == 3 {
		color = "Red"
	} else {
		color = "Out of range"
	}
	return color
}
