package vehicles

type Vehicle interface {
	Start() int
	Drive() int
	SpeedUp() int
	SpeedDown() int
	Stop() int
}
