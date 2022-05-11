package vehicles

import (
	"log"

	helper "github.com/CK-Gnap/traffic/helpers"
)

type Car struct {
	Brand  string
	Model  string
	Year   int
	Wheels int
	Color  string
	Speed  int
}

func (c *Car) Start() int {
	c.Speed += helper.RandInt(15, 30)
	log.Printf("Car : %s %s is started at %v km/h\n", c.Brand, c.Model, c.Speed)
	return c.Speed
}

func (c *Car) Drive() int {
	c.Speed += helper.RandInt(1, c.Speed)
	log.Printf("Car : %s %s is driving at %v km/h\n", c.Brand, c.Model, c.Speed)
	return c.Speed
}

func (c *Car) SpeedUp() int {
	c.Speed += helper.RandInt(1, c.Speed)
	log.Printf("Car : %s %s is speed up to %v km/h\n", c.Brand, c.Model, c.Speed)
	return c.Speed
}

func (c *Car) SpeedDown() int {
	c.Speed -= helper.RandInt(1, c.Speed)
	log.Printf("Car : %s %s is speed down to %v km/h\n", c.Brand, c.Model, c.Speed)
	return c.Speed
}

func (c *Car) Stop() int {
	c.Speed = helper.Minus(c.Speed, c.Speed)
	log.Printf("Car : %s %s is stopped at %v km/h\n", c.Brand, c.Model, c.Speed)
	return c.Speed
}
