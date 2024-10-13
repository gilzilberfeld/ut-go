package demos

type Driver struct {
	car ICar
}

func (d *Driver) CanStartDriving() bool {
	return !d.car.IsRunning()
}

func (d *Driver) Drive() {
	d.car.SetAC(AirCondition{mode: On})
	d.car.Start()
}