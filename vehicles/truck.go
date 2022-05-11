package vehicles

import (
	"log"

	helper "github.com/CK-Gnap/traffic/helpers"
)

type Truck struct {
	Brand  string
	Model  string
	Year   int
	Wheels int
	Color  string
	Speed  int
}

func (t *Truck) Start() int {
	t.Speed = helper.RandInt(20, 40)
	log.Printf("Truck : %s %s is started at %v km/h\n", t.Brand, t.Model, t.Speed)
	return t.Speed
}

func (t *Truck) Drive() int {
	t.Speed += helper.RandInt(1, t.Speed)
	log.Printf("Truck : %s %s is driving at %v km/h\n", t.Brand, t.Model, t.Speed)
	return t.Speed
}

func (t *Truck) SpeedUp() int {
	t.Speed += helper.RandInt(1, t.Speed)
	log.Printf("Truck : %s %s is speed up to %v km/h\n", t.Brand, t.Model, t.Speed)
	return t.Speed
}

func (t *Truck) SpeedDown() int {
	t.Speed -= helper.RandInt(1, t.Speed)
	log.Printf("Truck : %s %s is speed down to %v km/h\n", t.Brand, t.Model, t.Speed)
	return t.Speed
}

func (t *Truck) Stop() int {
	t.Speed = 0
	log.Printf("Truck : %s %s is stopped at %v km/h\n", t.Brand, t.Model, t.Speed)
	return t.Speed
}
