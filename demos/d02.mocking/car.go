package demos

type ICar interface {
	IsRunning() bool
	SetAC(ac AirCondition)
	Start()
}
type Car struct{}

func (car *Car) IsRunning() bool {
	panic("Not yet implemented")
}

func (car *Car) SetAC(ac AirCondition) {
	panic("Not yet implemented")
}

func (car *Car) Start() {
	panic("Not yet implemented")
}

func (car *Car) GetAC() AirCondition {
	panic("Not yet implemented")
}