package vehicles

import (
	"log"

	helper "github.com/CK-Gnap/traffic/helpers"
)

type Motorcycle struct {
	Brand  string
	Model  string
	Year   int
	Wheels int
	Color  string
	Speed  int
}

func (m *Motorcycle) Start() int {
	m.Speed += helper.RandInt(10, 20)
	log.Printf("Motorcycle : %s %s is started at %v km/h\n", m.Brand, m.Model, m.Speed)
	return m.Speed
}

func (m *Motorcycle) Drive() int {
	m.Speed += helper.RandInt(1, m.Speed)
	log.Printf("Motorcycle : %s %s is driving at %v km/h\n", m.Brand, m.Model, m.Speed)
	return m.Speed
}

func (m *Motorcycle) SpeedUp() int {
	m.Speed += helper.RandInt(1, m.Speed)
	log.Printf("Motorcycle : %s %s is speed up to %v km/h\n", m.Brand, m.Model, m.Speed)
	return m.Speed
}

func (m *Motorcycle) SpeedDown() int {
	m.Speed -= helper.RandInt(1, m.Speed)
	log.Printf("Motorcycle : %s %s is speed down to %v km/h\n", m.Brand, m.Model, m.Speed)
	return m.Speed
}

func (m *Motorcycle) Stop() int {
	m.Speed = 0
	log.Printf("Motorcycle : %s %s is stopped at %v km/h\n", m.Brand, m.Model, m.Speed)
	return m.Speed
}
